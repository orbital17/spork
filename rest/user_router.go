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
	router.HandleFunc("/me", Auth(s.handleMe()))
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

func (s *AccountRouter) handleMe() http.HandlerFunc {
	type response struct {
		Id   int64  `json:"id"`
		Test string `json:"test"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := GetUserID(r)
		user, err := s.service.GetById(id)
		if !checkInternalError(w, err) {
			return
		}
		WriteJson(w, &response{user.Id, "jflkds"})
	}
}
