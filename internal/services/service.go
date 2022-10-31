package services

import "test-task/internal/repository"

type Repository interface {
	AddTransaction(transaction *repository.Transaction) error
	GetByTransactionID(id string) (*repository.Transaction, error)
	GetByTerminalID(id string) (*repository.Transaction, error)
	GetByStatus(status string) ([]*repository.Transaction, error)
	GetByPaymentType(status string) ([]*repository.Transaction, error)
	GetByPeriod(day repository.Date) ([]*repository.Transaction, error)
	GetByNarrative(text string) ([]*repository.Transaction, error)
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

func (s *Service) GetByTransactionID(id string) (*repository.Transaction, error) {
	return s.repo.GetByTransactionID(id)
}

func (s *Service) GetByTerminalID(id string) (*repository.Transaction, error) {
	return s.repo.GetByTerminalID(id)
}

func (s *Service) GetByStatus(status string) ([]*repository.Transaction, error) {
	return s.repo.GetByStatus(status)
}

func (s *Service) GetByPaymentType(payment string) ([]*repository.Transaction, error) {
	return s.repo.GetByPaymentType(payment)
}

func (s *Service) GetByPeriod(day repository.Date) ([]*repository.Transaction, error) {
	return s.repo.GetByPeriod(day)
}

func (s *Service) GetByNarrative(text string) ([]*repository.Transaction, error) {
	return s.repo.GetByNarrative(text)
}
