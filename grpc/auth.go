package grpc_api

import (
	context "context"
	"spork/users"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

func auth(ctx context.Context) (auth *users.Auth, err error) {
	err = status.Errorf(codes.Unauthenticated, "auth failed")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}
	arr := md.Get("jwt")
	if len(arr) != 1 {
		return
	}
	token := arr[0]
	auth, parseErr := users.ParseToken(token)
	if parseErr != nil {
		return
	}
	return auth, nil
}

type serverOptionalAuth interface {
	OptionalAuth(fullMethod string) bool
}

func authUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	auth, err := auth(ctx)
	if err != nil {
		if optServer, ok := info.Server.(serverOptionalAuth); ok {
			if optServer.OptionalAuth(info.FullMethod) {
				return handler(ctx, req)
			}
		}
		return nil, err
	}
	newCtx := users.NewContext(ctx, auth)
	return handler(newCtx, req)
}

func withAuthInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(authUnaryInterceptor)
}
