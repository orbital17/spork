package grpc_api

import (
	context "context"
	"spork/users"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Server struct {
	UserService *users.Service
}

func NewServer(userService *users.Service) *Server {
	return &Server{userService}
}

func (*Server) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*Server) FindByEmail(ctx context.Context, req *FindByEmailRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
}
