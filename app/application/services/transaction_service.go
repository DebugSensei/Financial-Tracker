package services

import (
	"financial-tracker/interfaces/repository"
	"financial-tracker/internal/domain/transaction"
	"time"
)

type TransactionService struct {
	TransactionRepo repository.TransactionRepository
}

func (s *TransactionService) AddTransaction(amount float64, currency, tType string, categoryID int) (transaction.Transaction, error) {
	trans := transaction.Transaction{
		Date:       time.Now(),
		Amount:     amount,
		Currency:   currency,
		Type:       tType,
		CategoryID: categoryID,
	}
	err := s.TransactionRepo.AddTransaction(trans)
	return trans, err
}

func (s *TransactionService) GetBalance() (float64, error) {
	return s.TransactionRepo.GetBalance()
}

func (s *TransactionService) GetTransactions() ([]transaction.TransactionWithCategory, error) {
	return s.TransactionRepo.GetTransactions()
}
