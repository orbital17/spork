package rest

import (
	"net/http"
	"spork/users"
)

type UserRouter struct {
	service *users.Service
}

func (s *UserRouter) routes() {
	http.HandleFunc("/api/login", s.handleLogin())
}

func (s *UserRouter) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
