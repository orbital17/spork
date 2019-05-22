package rest

import (
	"net/http"
	"spork/account"
)

type AccountRouter struct {
	service *account.Service
}

func (s *AccountRouter) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/login", s.handleLogin())
	return router
}

func (s *AccountRouter) handleLogin() http.HandlerFunc {
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
