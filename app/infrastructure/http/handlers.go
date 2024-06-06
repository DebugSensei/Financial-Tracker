package http

import (
	"financial_tracker/application/services"
	"financial_tracker/interfaces/repository"
	"financial_tracker/internal/domain/category"
	"financial_tracker/internal/domain/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTransaction(c *gin.Context) {
	var input transaction.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionRepo := repository.NewTransactionRepository()
	service := services.NewTransactionService(transactionRepo)

	if err := service.AddTransaction(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction added successfully!"})
}

func GetBalance(c *gin.Context) {
	transactionRepo := repository.NewTransactionRepository()
	service := services.NewTransactionService(transactionRepo)

	balance, err := service.GetBalance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func AddCategory(c *gin.Context) {
	var input category.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryRepo := repository.NewCategoryRepository()
	service := services.NewCategoryService(categoryRepo)

	if err := service.AddCategory(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category added successfully!"})
}

func GetCategories(c *gin.Context) {
	categoryRepo := repository.NewCategoryRepository()
	service := services.NewCategoryService(categoryRepo)

	categories, err := service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetTransactions(c *gin.Context) {
	transactionRepo := repository.NewTransactionRepository()
	service := services.NewTransactionService(transactionRepo)

	transactions, err := service.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
