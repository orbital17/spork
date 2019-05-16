package grpc_api

import (
	fmt "fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type Runner struct {
	server *grpc.Server
}

func NewRunner(users *Users) *Runner {
	grpcServer := grpc.NewServer()
	RegisterUsersServer(grpcServer, users)
	return &Runner{grpcServer}
}

func (runner *Runner) Run() {
	port := 3000
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
