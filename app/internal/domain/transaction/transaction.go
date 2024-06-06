package transaction

import "time"

type Transaction struct {
	ID         int
	Date       time.Time
	Amount     float64
	Currency   string
	Type       string
	CategoryID int
}

type TransactionWithCategory struct {
	ID       int
	Date     time.Time
	Amount   float64
	Currency string
	Type     string
	Category string
}
