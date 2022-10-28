package main

import (
	"log"
	"test-task/internal/repository"
	"test-task/internal/server"
)

func main() {
	dbStr := "postgres://user:password@localhost:5432/test?sslmode=disable"
	db, err := repository.NewConnection(dbStr)

	if err != nil {
		log.Fatalf("%v", err)
	}

	_ = repository.NewRepository(db)
	s := server.NewServer()
	s.RunServer()
}
