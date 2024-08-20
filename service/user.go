package service

import (
	"context"

	pb "github.com/mirjalilova/auth-service-blacklist/internal/genproto/auth"
	st "github.com/mirjalilova/auth-service-blacklist/pkg/storage/postgres"
	"golang.org/x/exp/slog"
)

type UserService struct {
	storage st.Storage
	pb.UnimplementedUserServiceServer
}

func NewUserService(storage *st.Storage) *UserService {
	return &UserService{
		storage: *storage,
	}
}

func (s *UserService) GetProfile(ctx context.Context, req *pb.GetById) (*pb.UserRes, error) {
	res, err := s.storage.UserS.GetProfile(req)
	if err != nil {
        slog.Error("failed to get user profile: %v", err)
		return nil, err
	}

    slog.Info("got user profile: %+v", res)
	return res, nil
}

func (s *UserService) EditProfile(ctx context.Context, req *pb.UserRes) (*pb.UserRes, error) {
	res, err := s.storage.UserS.EditProfile(req)
    if err != nil {
        slog.Error("failed to edit user profile: %v", err)
        return nil, err
    }

    slog.Info("edited user profile: %+v", res)
    return res, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Void, error) {
	res, err := s.storage.UserS.ChangePassword(req)
    if err!= nil {
        slog.Error("failed to change user password: %v", err)
        return nil, err
    }

    slog.Info("changed user password: %+v", res)
    return res, nil
}

func (s *UserService) GetSetting(ctx context.Context, req *pb.GetById) (*pb.Setting, error) {
	res, err := s.storage.UserS.GetSetting(req)
    if err!= nil {
        slog.Error("failed to get user setting: %v", err)
        return nil, err
    }

    slog.Info("got user setting: %+v", res)
    return res, nil
}

func (s *UserService) EditSetting(ctx context.Context, req *pb.SettingReq) (*pb.Void, error) {
	res, err := s.storage.UserS.EditSetting(req)
    if err!= nil {
        slog.Error("failed to edit user setting: %v", err)
        return nil, err
    }

    slog.Info("edited user setting: %+v", res)
    return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.UserS.DeleteUser(req)
    if err!= nil {
        slog.Error("failed to delete user: %v", err)
        return nil, err
    }

    slog.Info("deleted user: %+v", res)
    return res, nil
}

