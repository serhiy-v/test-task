package services

import "test-task/internal/repository"

type Repository interface {
	AddTransaction(transaction *repository.Transaction) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddTransaction(transaction *repository.Transaction) error {
	return s.repo.AddTransaction(transaction)
}
