package main

import (
	"log"
	"net"

	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"github.com/Jerinji2016/grpc-template/src/pkg/middleware"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"github.com/Jerinji2016/grpc-template/src/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	logger.InitLogger()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	authInterceptor := middleware.NewAuthInterceptor()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)

	authService := service.NewAuthenticationService()
	postsService := service.NewPostServce()
	
	pb.RegisterAuthenticationServiceServer(grpcServer, authService)
	pb.RegisterPostServiceServer(grpcServer, postsService)

	log.Println("Serving at grpc://localhost:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
