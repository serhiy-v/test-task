package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RunServer() {
	router := NewRouter(s)
	addr := "localhost:8080"

	log.Printf("start HTTP server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func (s *Server) Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
