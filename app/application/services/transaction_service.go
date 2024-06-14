package services

import (
	"financial_tracker/app/internal/domain/transaction"
	"time"
)

type TransactionRepository interface {
	AddTransaction(trans transaction.Transaction) error
	GetBalance() (float64, error)
	GetTransactions() ([]transaction.TransactionWithCategory, error)
}

type TransactionService struct {
	TransactionRepo TransactionRepository
}

func NewTransactionService(repo TransactionRepository) *TransactionService {
	return &TransactionService{TransactionRepo: repo}
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
