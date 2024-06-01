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
