package service

import (
	"context"
	"fmt"

	"github.com/Jerinji2016/grpc-template/src/internal/auth"
	"github.com/Jerinji2016/grpc-template/src/internal/models"
	"github.com/Jerinji2016/grpc-template/src/internal/repositories"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthenticationService struct {
	pb.UnimplementedAuthenticationServiceServer
	userRepo *repositories.UserRepository
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (s *AuthenticationService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// hard-coded authentication check
	user, err := s.userRepo.FindUserByUsername(req.Username)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "user does not exist")
	}

	if req.Password != user.Password {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (s *AuthenticationService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to create user: %v", err))
	}

	return &pb.RegisterResponse{
		Id:       user.ID.UUID.String(),
		Name:     user.Name,
		Username: user.Username,
	}, nil
}
