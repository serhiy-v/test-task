package main

import (
	"log"
	"test-task/internal/repository"
	"test-task/internal/server"
	"test-task/internal/services"
)

func main() {
	dbStr := "postgres://user:password@localhost:5432/test?sslmode=disable"
	db, err := repository.NewConnection(dbStr)

	if err != nil {
		log.Fatalf("%v", err)
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	s := server.NewServer(service)
	s.RunServer()
}
