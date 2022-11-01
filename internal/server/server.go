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
			log.Println(err)
			return
		}
	}
}

func (s *Server) GetByTransactionID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	trx, err := s.service.GetByTransactionID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(trx)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) GetByTerminalID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ids := strings.Split(id, ",")

	for _, id := range ids {
		trx, err := s.service.GetByTransactionID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = json.NewEncoder(w).Encode(trx)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) GetByStatus(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]

	transactions, err := s.service.GetByStatus(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	for _, transaction := range transactions {
		err = json.NewEncoder(w).Encode(transaction)
		if err != nil {
			log.Printf("Unable to encode transaction: %v\n", err)
			return
		}
	}
}

func (s *Server) GetByPaymentType(w http.ResponseWriter, r *http.Request) {
	payment := mux.Vars(r)["payment"]

	transactions, err := s.service.GetByPaymentType(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	for _, transaction := range transactions {
		err = json.NewEncoder(w).Encode(transaction)
		if err != nil {
			log.Printf("Unable to encode transaction: %v\n", err)
			return
		}
	}
}

func (s *Server) GetForPeriod(w http.ResponseWriter, r *http.Request) {
	var date repository.Date
	err := json.NewDecoder(r.Body).Decode(&date)
	if err != nil {
		log.Printf("Unable to get date from request: %v\n", err)
	}

	transactions, err := s.service.GetByPeriod(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	for _, transaction := range transactions {
		err = json.NewEncoder(w).Encode(transaction)
		if err != nil {
			log.Printf("Unable to encode transaction: %v\n", err)
			return
		}
	}
}

func (s *Server) GetByNarrative(w http.ResponseWriter, r *http.Request) {
	var text repository.NarrativeText
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		log.Printf("Unable to get date from request: %v\n", err)
	}

	transactions, err := s.service.GetByNarrative(text.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	for _, transaction := range transactions {
		err = json.NewEncoder(w).Encode(transaction)
		if err != nil {
			log.Printf("Unable to encode transaction: %v\n", err)
			return
		}
	}
}
