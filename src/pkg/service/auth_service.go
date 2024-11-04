package service

import (
	"context"

	"github.com/Jerinji2016/grpc-template/src/pkg/auth"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthenticationService struct {
	pb.UnimplementedAuthenticationServiceServer
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (s *AuthenticationService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// hard-coded authentication check
	if req.Username != "admin" || req.Password != "secret" {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &pb.LoginResponse{Token: token}, nil
}
