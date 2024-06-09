package repository

import (
	"database/sql"
	"financial_tracker/app/internal/domain/transaction"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) AddTransaction(trans transaction.Transaction) error {
	query := `INSERT INTO transactions (date, amount, currency, type, category_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, trans.Date, trans.Amount, trans.Currency, trans.Type, trans.CategoryID)
	return err
}

func (r *TransactionRepository) GetBalance() (float64, error) {
	var balance float64
	query := `SELECT SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END) as balance FROM transactions`
	err := r.db.QueryRow(query).Scan(&balance)
	return balance, err
}

func (r *TransactionRepository) GetTransactions() ([]transaction.TransactionWithCategory, error) {
	var transactions []transaction.TransactionWithCategory
	query := `
		SELECT t.date, t.amount, t.currency, t.type, c.name as category_name
		FROM transactions t
		JOIN categories c ON t.category_id = c.id
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trans transaction.TransactionWithCategory
		if err := rows.Scan(&trans.Date, &trans.Amount, &trans.Currency, &trans.Type, &trans.CategoryName); err != nil {
			return nil, err
		}
		transactions = append(transactions, trans)
	}

	return transactions, nil
}
