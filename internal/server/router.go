package server

import (
	"github.com/gorilla/mux"
)

func NewRouter(s *Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.Greeting).Methods("GET")
	r.HandleFunc("/", s.UploadFile).Methods("POST")

	return r
}
