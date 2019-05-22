package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"spork/auth"
	"strings"
)

func checkJsonError(w http.ResponseWriter, err error) (ok bool) {
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, "Bad request: can't parse json")
		return false
	}
	return true
}

func checkInternalError(w http.ResponseWriter, err error) (ok bool) {
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err.Error())
		log.Print(err)
		return false
	}
	return true
}

func DecodeJson(w http.ResponseWriter, r *http.Request, req interface{}) (ok bool) {
	err := json.NewDecoder(r.Body).Decode(req)
	return checkJsonError(w, err)
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if !checkInternalError(w, err) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if !checkInternalError(w, err) {
		return
	}
}

func writeError(w http.ResponseWriter, err string) {
	js, _ := json.Marshal(map[string]string{
		"error": err,
	})
	_, _ = w.Write(js)
}

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("authorization")
		if len(header) < 8 || strings.ToLower(header[0:7]) != "bearer " {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := header[7:]
		authObj, err := auth.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		newContext := auth.NewContext(r.Context(), authObj)
		newRequest := r.WithContext(newContext)
		handler(w, newRequest)
	}
}

func GetUserID(r *http.Request) int64 {
	authObj, _ := auth.FromContext(r.Context())
	return authObj.UserID
}
