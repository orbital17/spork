package rest

import (
	"encoding/json"
	"fmt"
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
	type request struct {
		Email    string
		Password string
	}
	type response struct {
		Token string
	}
	f := func(req *request) (res *response, err error) {
		token, err := s.service.Login(req.Email, req.Password)
		return &response{token}, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			fmt.Print(err)
			return
		}
		res, err := f(req)
		if err != nil {
			fmt.Print(err)
			return
		}
		js, err := json.Marshal(res)
		if err != nil {
			fmt.Print(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(js)
		if err != nil {
			fmt.Print(err)
			return
		}
	}
}
