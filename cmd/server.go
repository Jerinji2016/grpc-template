package main

import (
	"log"
	"net"

	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"github.com/Jerinji2016/grpc-template/src/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	postsService := service.NewPostServce()

	pb.RegisterPostServiceServer(grpcServer, postsService)

	log.Println("Serving at grpc://localhost:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
