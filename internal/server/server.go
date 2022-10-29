package server

import (
	"log"
	"net/http"
	"test-task/internal/parser"
	"test-task/internal/services"
)

type Server struct {
	service *services.Service
}

func NewServer(service *services.Service) *Server {
	return &Server{service: service}
}

func (s *Server) RunServer() {
	router := NewRouter(s)
	addr := "localhost:8080"

	log.Printf("start HTTP server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30 << 22)

	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	transactions := parser.GetData(file, header)

	for _, trs := range transactions {
		s.service.AddTransaction(trs)
	}
}
