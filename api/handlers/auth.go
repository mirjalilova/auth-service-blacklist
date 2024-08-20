package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/go-redis/redis"
	t "github.com/mirjalilova/auth-service-blacklist/api/token"
	"github.com/mirjalilova/auth-service-blacklist/internal/genproto/auth"
	"golang.org/x/exp/slog"

	"github.com/gin-gonic/gin"
	md "github.com/mirjalilova/auth-service-blacklist/api/middleware"
	"github.com/mirjalilova/auth-service-blacklist/pkg/email"
)

var (
	registeredUsers = struct {
		sync.RWMutex
		users map[string]struct{}
	}{
		users: make(map[string]struct{}),
	}

	from     = "mirjalilovaferuza952@gmail.com"
	password = "icuu qtla cwlc wstr"
)

// RegisterUser handles user registration
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body auth.RegisterReq true "Register User Request"
// @Success 200 {object} string "User registered successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /register [post]
func (h *Handlers) RegisterUser(c *gin.Context) {
	var body auth.RegisterReq
	if err := c.BindJSON(&body); err != nil {
		slog.Error("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	// Check if the username or email is already registered
	registeredUsers.RLock()
	if _, exists := registeredUsers.users[body.Email]; exists {
		registeredUsers.RUnlock()
		slog.Error("email already registered")
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already registered"})
		return
	}

	if _, exists := registeredUsers.users[body.Username]; exists {
		registeredUsers.RUnlock()
		slog.Error("username already taken")
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already taken"})
		return
	}
	registeredUsers.RUnlock()

	password, err := t.HashPassword(body.Password)
	if err != nil {
		slog.Error("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	req := &auth.RegisterReq{
		Username:    body.Username,
		Email:       body.Email,
		Password:    password,
		FullName:    body.FullName,
		DateOfBirth: body.DateOfBirth,
	}

	if !isValidEmail(req.Email) {
		slog.Error("invalid email format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
		return
	}

	if _, err := time.Parse("2006-01-02", req.DateOfBirth); err != nil {
		slog.Error("invalid date of birth format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date of birth format"})
		return
	}

	input, err := json.Marshal(req)
	if err != nil {
		slog.Error("failed to marshal JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = h.Producer.ProduceMessages("reg-user", input)
	if err != nil {
		slog.Error("failed to produce message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Add the new user to the map
	registeredUsers.Lock()
	registeredUsers.users[req.Email] = struct{}{}
	registeredUsers.users[req.Username] = struct{}{}
	registeredUsers.Unlock()

	slog.Info("User registered successfully", "username", req.Username)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body auth.LoginReq true "Login Request"
// @Success 200 {string} auth.LoginRes
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /login [post]
func (h *Handlers) LoginUser(c *gin.Context) {
	var req auth.LoginReq
	if err := c.BindJSON(&req); err != nil {
		slog.Error("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	res, err := h.Auth.Login(context.Background(), &req)
	if err != nil {
		slog.Error("failed to login user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	token, refToken := t.GenerateJWTToken(res)

	slog.Info("User logged in successfully", "username", req.Username)
	c.JSON(http.StatusOK, auth.LoginRes{
		AccessToken:  token,
		RefreshToken: refToken,
	})
}

// ForgotPassword handles forgot password functionality
// @Summary Forgot password
// @Description Request to reset user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.GetByEmail true "Email Request"
// @Success 200 {string} string "Password reset email sent successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /forgot-password [post]
func (h *Handlers) ForgotPassword(c *gin.Context) {
	var req auth.GetByEmail
	if err := c.BindJSON(&req); err != nil {
		slog.Error("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	_, err := h.Auth.ForgotPassword(context.Background(), &req)
	if err != nil {
		slog.Error("failed to send password reset email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	forgotPasswordCode := email.GenForgotPassword()

	err = h.RDB.Set(context.Background(), forgotPasswordCode, req.Email, 15*time.Minute).Err()
	if err != nil {
		slog.Error("failed to store forgot password code in Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = email.SendVerificationCode(&auth.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi %s, your verification:%s", req.Email, forgotPasswordCode),
		Code:     forgotPasswordCode,
	})

	if err != nil {
		slog.Error("Could not send an email: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Password reset email sent successfully", "email", req.Email)
	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}

// ResetPassword handles password reset
// @Summary Reset password
// @Description Reset user's password with a reset code
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.ResetPassReqBody true "Password Reset Request"
// @Success 200 {string} string "Password reset successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /reset-password [post]
func (h *Handlers) ResetPassword(c *gin.Context) {
	var req auth.ResetPassReq
	if err := c.BindJSON(&req); err != nil {
		slog.Error("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	password, err := t.HashPassword(req.NewPassword)
	if err != nil {
		slog.Error("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	req.NewPassword = password

	email, err := h.RDB.Get(context.Background(), req.ResetToken).Result()
	if err == redis.Nil {
		slog.Error("forgot password code not found in Redis: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.Email = email

	_, err = h.Auth.ResetPassword(context.Background(), &req)
	if err != nil {
		slog.Error("failed to reset password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Password reset successfully", "email", email)
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

func isValidEmail(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

func getuserId(ctx *gin.Context) string {
	user_id, err := md.GetUserId(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}
	return user_id
}
