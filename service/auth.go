package service

import (
	"context"

	pb "github.com/mirjalilova/auth-service-blacklist/internal/genproto/auth"
	st "github.com/mirjalilova/auth-service-blacklist/pkg/storage/postgres"
	"golang.org/x/exp/slog"
)

type AuthService struct {
	storage st.Storage
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(storage *st.Storage) *AuthService {
	return &AuthService{
		storage: *storage,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.Void, error) {
	res, err := s.storage.AuthS.Register(req)
	if err != nil {
		slog.Error("Error registering user: %v", err)
		return nil, err
	}

	slog.Info("User registered successfully: %s", res)
	return res, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.User, error) {
	res, err := s.storage.AuthS.Login(req)
	if err != nil {
		slog.Error("Error logging in user: %v", err)
		return nil, err
	}

	slog.Info("User logged in successfully: %s", res)
	return res, nil
}


func (s *AuthService) ForgotPassword(ctx context.Context, req *pb.GetByEmail) (*pb.Void, error) {
	res, err := s.storage.AuthS.ForgotPassword(req)
    if err != nil {
		slog.Error("Error sending forgot password email: %v", err)
        return nil, err
    }

	slog.Info("Forgot password email sent successfully: %s", req.Email)
    return res, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPassReq) (*pb.Void, error) {
	res, err := s.storage.AuthS.ResetPassword(req)
    if err != nil {
		slog.Error("Error resetting password: %v", err)
        return nil, err
    }

	slog.Info("Password reset successfully: %s", req.Email)
    return res, nil
}

func (s *AuthService) SaveRefreshToken(ctx context.Context, req *pb.RefToken) (*pb.Void, error) {
	res, err := s.storage.AuthS.SaveRefreshToken(req)
    if err != nil {
		slog.Error("Error saving refresh token: %v", err)
        return nil, err
    }

	slog.Info("Refresh token saved successfully")
    return res, nil
}

func (s *AuthService) GetAllUsers(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserRes, error) {
	res, err := s.storage.AuthS.GetAllUsers(req)
    if err!= nil {
        slog.Error("Error getting all users: %v", err)
        return nil, err
    }

    slog.Info("Got all users: %+v", res)
    return res, nil
}