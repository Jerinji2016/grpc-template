package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"github.com/Jerinji2016/grpc-template/src/pkg/middleware"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"github.com/Jerinji2016/grpc-template/src/pkg/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	logger.InitLogger()

	if err := godotenv.Load(); err != nil {
		logger.FatalLog("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")
	address := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.FatalLog("Failed to listen: %v", err)
	}

	authInterceptor := middleware.NewAuthInterceptor()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)

	authService := service.NewAuthenticationService()
	postsService := service.NewPostServce()

	pb.RegisterAuthenticationServiceServer(grpcServer, authService)
	pb.RegisterPostServiceServer(grpcServer, postsService)

	logger.InfoLog("Serving at grpc://localhost%s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
