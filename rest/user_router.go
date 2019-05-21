package rest

import (
	"net/http"
	"spork/users"
)

type UserRouter struct {
	service *users.Service
}

func (s *UserRouter) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/login", s.handleLogin())
	return router
}

func (s *UserRouter) handleLogin() http.HandlerFunc {
	type request struct {
		Email    string
		Password string
	}
	type response struct {
		Token string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if !DecodeJson(w, r, req) {
			return
		}
		token, err := s.service.Login(req.Email, req.Password)
		if !checkInternalError(w, err) {
			return
		}
		WriteJson(w, &response{token})
	}
}
