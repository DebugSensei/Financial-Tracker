package api

import (
	"financial_tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Temporary storage for transactions
var transactions []models.Transaction

// AddTransaction handles adding a new transaction
func AddTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction.Date = time.Now()
	// Add transaction to storage
	transactions = append(transactions, transaction)
	c.JSON(http.StatusOK, transaction)
}

// GetBalance handles fetching the balance at a specific time or current time
func GetBalance(c *gin.Context) {
	balance := 0.0
	for _, transaction := range transactions {
		if transaction.Type == "income" {
			balance += transaction.Amount
		} else if transaction.Type == "expense" {
			balance -= transaction.Amount
		}
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
