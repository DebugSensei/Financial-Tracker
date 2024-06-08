package services

import (
	"financial_tracker/app/interfaces/repository"
	"financial_tracker/app/internal/domain/transaction"
)

// TransactionService struct
type TransactionService struct {
	TransactionRepo repository.TransactionRepository
}

// NewTransactionService creates a new TransactionService
func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{TransactionRepo: repo}
}

func (s *TransactionService) AddTransaction(trans transaction.Transaction) error {
	return s.TransactionRepo.AddTransaction(trans)
}

func (s *TransactionService) GetBalance() (float64, error) {
	return s.TransactionRepo.GetBalance()
}

func (s *TransactionService) GetTransactions() ([]transaction.TransactionWithCategory, error) {
	return s.TransactionRepo.GetTransactions()
}
