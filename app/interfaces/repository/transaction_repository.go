package repository

import (
	"database/sql"
	"financial-tracker/internal/domain/transaction"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{DB: db}
}

func (r *TransactionRepository) AddTransaction(trans transaction.Transaction) error {
	query := `INSERT INTO transactions (date, amount, currency, type, category_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, trans.Date, trans.Amount, trans.Currency, trans.Type, trans.CategoryID)
	return err
}

func (r *TransactionRepository) GetBalance() (float64, error) {
	var balance float64
	query := `SELECT COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END), 0) FROM transactions`
	err := r.DB.QueryRow(query).Scan(&balance)
	return balance, err
}

func (r *TransactionRepository) GetTransactions() ([]transaction.TransactionWithCategory, error) {
	query := `
        SELECT t.id, t.date, t.amount, t.currency, t.type, c.name
        FROM transactions t
        JOIN categories c ON t.category_id = c.id
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []transaction.TransactionWithCategory
	for rows.Next() {
		var trans transaction.TransactionWithCategory
		if err := rows.Scan(&trans.ID, &trans.Date, &trans.Amount, &trans.Currency, &trans.Type, &trans.Category); err != nil {
			return nil, err
		}
		transactions = append(transactions, trans)
	}
	return transactions, nil
}
