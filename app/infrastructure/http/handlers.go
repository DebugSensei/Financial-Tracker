package http

import (
	"database/sql"
	"financial_tracker/app/application/services"
	"financial_tracker/app/interfaces/repository"
	"financial_tracker/app/internal/domain/category"
	"financial_tracker/app/internal/domain/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

type Handler struct {
	DB *sql.DB
}

func (h *Handler) AddTransaction(c *gin.Context) {
	var input transaction.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionRepo := repository.NewTransactionRepository(h.DB)
	service := services.NewTransactionService(transactionRepo)

	trans, err := service.AddTransaction(input.Amount, input.Currency, input.Type, input.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction added successfully!", "transaction": trans})
}

func (h *Handler) GetBalance(c *gin.Context) {
	transactionRepo := repository.NewTransactionRepository(h.DB)
	service := services.NewTransactionService(transactionRepo)

	balance, err := service.GetBalance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (h *Handler) AddCategory(c *gin.Context) {
	var input category.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryRepo := repository.NewCategoryRepository(h.DB)
	service := services.NewCategoryService(categoryRepo)

	cat, err := service.AddCategory(input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category added successfully!", "category": cat})
}

func (h *Handler) GetCategories(c *gin.Context) {
	categoryRepo := repository.NewCategoryRepository(h.DB)
	service := services.NewCategoryService(categoryRepo)

	categories, err := service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	transactionRepo := repository.NewTransactionRepository(h.DB)
	service := services.NewTransactionService(transactionRepo)

	transactions, err := service.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/transactions", h.AddTransaction)
	router.GET("/balance", h.GetBalance)
	router.POST("/categories", h.AddCategory)
	router.GET("/categories", h.GetCategories)
	router.GET("/transactions", h.GetTransactions)

	return router
}
