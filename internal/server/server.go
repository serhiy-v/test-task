package server

import (
	"fmt"
	"log"
	"net/http"
	"test-task/internal/parser"
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

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30 << 22)

	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	transactions := parser.GetData(file, header)

	for _, trs := range transactions {
		fmt.Println(trs)
	}
	return
}
