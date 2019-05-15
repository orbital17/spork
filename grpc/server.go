package grpc_api

import (
	context "context"
	"spork/users"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Server struct {
	UserService *users.Service
	UserStore   *users.Store
}

func NewServer(userService *users.Service, userStore *users.Store) *Server {
	return &Server{userService, userStore}
}

func (s *Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	token, err := s.UserService.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Token: token}, nil
}

func (*Server) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (s *Server) FindByEmail(ctx context.Context, req *FindByEmailRequest) (*User, error) {
	user, err := s.UserStore.UserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	res := &User{
		Id:    int64(user.Id),
		Email: user.Email,
		Name:  user.Name,
	}
	return res, nil
}
