package transaction

import (
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Currency   string    `json:"currency"`
	Type       string    `json:"type"`
	CategoryID int       `json:"category_id"`
}

type TransactionWithCategory struct {
	Date         time.Time `json:"date"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	Type         string    `json:"type"`
	CategoryName string    `json:"category_name"`
}
