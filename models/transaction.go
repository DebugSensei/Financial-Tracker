package models

import "time"

// Transaction represents a financial transaction
type Transaction struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Currency   string    `json:"currency"`
	Type       string    `json:"type"`
	CategoryID int       `json:"category_id"`
}

// TransactionWithCategory represents a transaction with category name
type TransactionWithCategory struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Type     string    `json:"type"`
	Category string    `json:"category"`
}
