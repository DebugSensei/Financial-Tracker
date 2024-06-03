package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/transaction", AddTransaction)
	r.GET("/balance", GetBalance)
	// r.POST("/category", AddCategory)
	r.GET("/categories", GetCategories)
	r.GET("/transactions", GetTransactions)
	return r
}
