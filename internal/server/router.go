package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(s *Server) *mux.Router {
	r := mux.NewRouter()
	r.Use(jsonHeader)

	r.HandleFunc("/", s.UploadFile).Methods("POST")
	r.HandleFunc("/transaction_id/{id}", s.GetByTransactionID).Methods("GET")
	r.HandleFunc("/terminal_id/{id}", s.GetByTerminalID).Methods("GET")
	r.HandleFunc("/status/{status}", s.GetByStatus).Methods("GET")
	r.HandleFunc("/payment_type/{payment}", s.GetByPaymentType).Methods("GET")
	r.HandleFunc("/period", s.GetForPeriod).Methods("GET")
	r.HandleFunc("/narrative", s.GetByNarrative).Methods("GET")

	return r
}

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		next.ServeHTTP(w, r)
	})
}
