package server

import (
	"github.com/gorilla/mux"
)

func NewRouter(s *Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.UploadFile).Methods("POST")
	r.HandleFunc("/transaction_id/{id}", s.GetByTransactionID).Methods("GET")
	r.HandleFunc("/terminal_id/{id}", s.GetByTerminalID).Methods("GET")
	r.HandleFunc("/status/{status}", s.GetByStatus).Methods("GET")
	r.HandleFunc("/payment_type/{payment}", s.GetByPaymentType).Methods("GET")
	r.HandleFunc("/period", s.GetForPeriod).Methods("GET")
	r.HandleFunc("/narrative", s.GetByNarrative).Methods("GET")

	return r
}
