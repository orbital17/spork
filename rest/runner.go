package rest

import (
	"fmt"
	"log"
	"net/http"
	"spork/config"
	"spork/users"
)

type Server struct {
	config     *config.Config
	userRouter *UserRouter
}

func NewServer(cfg *config.Config, users *users.Service) *Server {
	return &Server{
		cfg,
		&UserRouter{users},
	}
}

func (r *Server) Serve() {
	r.userRouter.routes()
	port := fmt.Sprintf(":%v", r.config.RestPort)
	fmt.Printf("rest listening to %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
