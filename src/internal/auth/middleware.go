package auth

import (
	"context"
	"slices"
	"strings"

	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type keyType string

const (
	CLAIMS_KEY        keyType = "claims-key"
	AUTHORIZATION_KEY string  = "authorization"
)

type AuthInterceptor struct{}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.DebugLog("Received request for %s", info.FullMethod)

		if a.isPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		claims, err := a.authorize(ctx)
		if err != nil {
			logger.ErrorLog("Authorization Failed %s, %v", info.FullMethod, err)
			return nil, err
		}

		ctx = context.WithValue(ctx, CLAIMS_KEY, claims)
		return handler(ctx, req)
	}
}

func (a *AuthInterceptor) isPublicMethod(fullMethod string) bool {
	var publicMethods = []string{
		pb.AuthenticationService_Login_FullMethodName,
		pb.AuthenticationService_Register_FullMethodName,
	}
	return slices.Contains(publicMethods, fullMethod)
}

func (a *AuthInterceptor) authorize(ctx context.Context) (jwt.MapClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	tokens := md[AUTHORIZATION_KEY]
	if len(tokens) == 0 || tokens[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "authorization token is missing")
	}

	tokenHeaders := tokens[0]
	token := strings.Split(tokenHeaders, " ")[1]

	claims, err := ValidateToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to extract claims")
	}

	return claims, nil
}
