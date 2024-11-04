package middleware

import (
	"context"
	"log"
	"strings"

	"github.com/Jerinji2016/grpc-template/src/internal/keys"
	"github.com/Jerinji2016/grpc-template/src/pkg/auth"
	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

		if a.isPublidMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		claims, err := a.authorize(ctx)
		if err != nil {
			logger.ErrorLog("Authorization Failed %s, %v", info.FullMethod, err)
			return nil, err
		}

		ctx = context.WithValue(ctx, keys.CLAIMS_KEY, claims)
		return handler(ctx, req)
	}
}

func (a *AuthInterceptor) isPublidMethod(fullMethod string) bool {
	// FullMethod is in the format "/package.Service/Method"
	log.Printf("full method %v", fullMethod)
	return strings.HasPrefix(fullMethod, "/auth_api.AuthenticationService/Login")
}

func (a *AuthInterceptor) authorize(ctx context.Context) (jwt.MapClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	tokens := md[keys.AUTHORIZATION_KEY]
	if len(tokens) == 0 || tokens[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "authorization token is missing")
	}

	tokenHeaders := tokens[0]
	token := strings.Split(tokenHeaders, " ")[1]

	claims, err := auth.ValidateToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to extract claims")
	}

	return claims, nil
}
