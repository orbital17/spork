package grpc_api

import (
	context "context"
	fmt "fmt"
	"spork/users"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Users struct {
	UserService *users.Service
	UserStore   *users.Store
}

func NewUsersServer(userService *users.Service, userStore *users.Store) *Users {
	return &Users{userService, userStore}
}

func (s *Users) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	token, err := s.UserService.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Token: token}, nil
}

func (*Users) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (s *Users) FindByEmail(ctx context.Context, req *FindByEmailRequest) (*User, error) {
	auth, _ := users.FromContext(ctx)
	fmt.Printf("id from context: %v", auth.UserID)
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
