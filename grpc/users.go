package grpc_api

import (
	context "context"
	"spork/users"

	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

func Auth(ctx context.Context) (userID users.UserID, err error) {
	err = status.Errorf(codes.Unauthenticated, "auth failed")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}
	arr := md.Get("authtoken")
	if len(arr) != 1 {
		return
	}
	token := arr[0]
	id, parseErr := users.ParseToken(token)
	if parseErr != nil {
		return
	}
	return id, nil
}

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
	_, err := Auth(ctx)
	if err != nil {
		return nil, err
	}
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
