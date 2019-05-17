package grpc_api

import (
	fmt "fmt"
	"log"
	"net"
	"spork/config"

	grpc "google.golang.org/grpc"
)

type Runner struct {
	server *grpc.Server
	config *config.Config
}

func NewRunner(users *Users, c *config.Config) *Runner {
	grpcServer := grpc.NewServer(
		withAuthInterceptor(),
	)
	RegisterUsersServer(grpcServer, users)
	return &Runner{grpcServer, c}
}

func (runner *Runner) Run() {
	port := runner.config.GrpcPort
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("listening to :%v\n", port)
	err = runner.server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to run the server: %v", err)
	}
}
