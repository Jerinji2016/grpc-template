package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Jerinji2016/grpc-template/src/internal/auth"
	"github.com/Jerinji2016/grpc-template/src/internal/db"
	"github.com/Jerinji2016/grpc-template/src/internal/service"
	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	logger.InitLogger()
}

func main() {
	db.InitDB()
	defer db.CloseDB()

	port := os.Getenv("PORT")
	address := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.FatalLog("Failed to listen: %v", err)
	}

	authInterceptor := auth.NewAuthInterceptor()
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
