package service

import (
	"Authorization_Service/internal/pb"
	"Authorization_Service/internal/postgres"
	"context"
)

type Server struct {
	H postgres.Handler
}

func (s *Server) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return &pb.SignUpResponse{
		Status: 0,
		Error:  "sadsad",
	}, nil
}

func (s *Server) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{}, nil
}
