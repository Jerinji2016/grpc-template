package service

import (
	"context"

	"github.com/Jerinji2016/grpc-template/src/internal/models"
	"github.com/Jerinji2016/grpc-template/src/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostService struct {
	pb.UnimplementedPostServiceServer
	posts map[string]*models.Post
}

func NewPostServce() *PostService {
	return &PostService{
		posts: make(map[string]*models.Post),
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := &models.Post{
		Message:   req.Post.Message,
		CreatedAt: req.Post.CreatedAt,
	}

	s.posts[post.ID.UUID.String()] = post

	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:        post.ID.UUID.String(),
			Message:   post.Message,
			CreatedAt: post.CreatedAt,
		},
	}, nil
}

func (s *PostService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, exists := s.posts[req.Id]
    if !exists {
        return nil, status.Error(codes.NotFound, "resource not found")
    }

    return &pb.GetPostResponse{
        Post: &pb.Post{
            Id:        post.ID.UUID.String(),
            Message:   post.Message,
            CreatedAt: post.CreatedAt,
        },
    }, nil
}