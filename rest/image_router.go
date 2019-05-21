package rest

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ImageRouter struct {
}

func (s *ImageRouter) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/upload", s.handleUpload())
	return router
}

func (s *ImageRouter) handleUpload() http.HandlerFunc {
	type response struct {
		FileId string `json:"fileId"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("image")
		if err != nil {
			writeError(w, err.Error())
			return
		}
		fileId := "downloaded"
		fileName := fmt.Sprintf("./temp/%v.jpg", fileId)
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			writeError(w, err.Error())
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			writeError(w, err.Error())
			return
		}
		WriteJson(w, &response{fileId})

	}
}
