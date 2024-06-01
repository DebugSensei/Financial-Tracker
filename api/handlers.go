package api

import (
	"financial_tracker/db"
	"financial_tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AddTransaction handles adding a new transaction
func AddTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction.Date = time.Now()
	query := `INSERT INTO transactions (date, amount, currency, type, category_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.DB.Exec(query, transaction.Date, transaction.Amount, transaction.Currency, transaction.Type, transaction.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// GetBalance handles fetching the balance at a specific time or current time
func GetBalance(c *gin.Context) {
	var balance float64
	query := `SELECT SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END) FROM transactions`
	err := db.DB.QueryRow(query).Scan(&balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// AddCategory handles adding a new category
func AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	err := db.DB.QueryRow(query, category.Name).Scan(&category.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// GetCategories handles fetching all categories
func GetCategories(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, name FROM categories`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetTransactions handles fetching all transactions with category names
func GetTransactions(c *gin.Context) {
	query := `
        SELECT t.id, t.date, t.amount, t.currency, t.type, c.name
        FROM transactions t
        JOIN categories c ON t.category_id = c.id
    `
	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []models.TransactionWithCategory
	for rows.Next() {
		var transaction models.TransactionWithCategory
		if err := rows.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.Currency, &transaction.Type, &transaction.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
