package transaction

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Currency   string    `json:"currency"`
	Type       string    `json:"type"`
	CategoryID int       `json:"category_id"`
}

type TransactionWithCategory struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Type     string    `json:"type"`
	Category string    `json:"category"`
}
