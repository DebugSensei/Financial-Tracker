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

// Handler struct holds required services for handler to function
type Handler struct {
	transactionService *services.TransactionService
	categoryService    *services.CategoryService
}

// NewHandler returns a new handler
func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		transactionService: services.NewTransactionService(repository.NewTransactionRepository(db)),
		categoryService:    services.NewCategoryService(repository.NewCategoryRepository(db)),
	}
}

func (h *Handler) AddTransaction(c *gin.Context) {
	var input transaction.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.transactionService.AddTransaction(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction added successfully!"})
}

func (h *Handler) GetBalance(c *gin.Context) {
	balance, err := h.transactionService.GetBalance()
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

	if err := h.categoryService.AddCategory(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category added successfully!"})
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	transactions, err := h.transactionService.GetTransactions()
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
