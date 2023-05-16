package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"test-task/internal/parser"
	"test-task/internal/repository"
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
	addr := "localhost:8090"

	log.Printf("start HTTP server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30 << 22)

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	transactions := parser.GetData(file, header)

	for _, trx := range transactions {
		err = s.service.AddTransaction(trx)
		if err != nil {
			log.Printf("Unable to add transaction: %v\n", err)
			return
		}
	}
}

func (s *Server) GetByTransactionID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	trx, err := s.service.GetByTransactionID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Unable to get transaction: %v\n", err)
		return
	}

	err = json.NewEncoder(w).Encode(trx)
	if err != nil {
		log.Printf("Unable to encode transaction: %v\n", err)
	}
}

func (s *Server) GetByTerminalID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ids := strings.Split(id, ",")

	transactions := make([]*repository.Transaction, 0)

	for _, id := range ids {
		trx, err := s.service.GetByTerminalID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Unable to get transaction: %v\n", err)
			return
		}

		transactions = append(transactions, trx)
	}

	err := json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Printf("Unable to encode transaction: %v\n", err)
	}
}

func (s *Server) GetByStatus(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]

	transactions, err := s.service.GetByStatus(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Unable to get transaction: %v\n", err)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Printf("Unable to encode transaction: %v\n", err)
		return
	}
}

func (s *Server) GetByPaymentType(w http.ResponseWriter, r *http.Request) {
	payment := mux.Vars(r)["payment"]

	transactions, err := s.service.GetByPaymentType(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Unable to get transaction: %v\n", err)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Printf("Unable to encode transaction: %v\n", err)
		return
	}
}

func (s *Server) GetForPeriod(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	transactions, err := s.service.GetByPeriod(from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Unable to get transaction: %v\n", err)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Printf("Unable to encode transaction: %v\n", err)
		return
	}
}

func (s *Server) GetByNarrative(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")

	transactions, err := s.service.GetByNarrative(text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Unable to get transaction: %v\n", err)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Printf("Unable to encode transaction: %v\n", err)
		return
	}
}
