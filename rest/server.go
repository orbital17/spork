package rest

import (
	"fmt"
	"log"
	"net/http"
	"spork/account"
	"spork/config"
)

type Server struct {
	config        *config.Config
	accountRouter *AccountRouter
}

func NewServer(cfg *config.Config, account *account.Service) *Server {
	return &Server{
		cfg,
		&AccountRouter{account},
	}
}

func (r *Server) Serve() {
	imageRouter := &ImageRouter{}
	router := http.NewServeMux()
	router.Handle("/api/account/", http.StripPrefix("/api/account", r.accountRouter.routes()))
	router.Handle("/api/images/", http.StripPrefix("/api/images", imageRouter.routes()))
	port := fmt.Sprintf(":%v", r.config.RestPort)
	fmt.Printf("rest listening to %v\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
