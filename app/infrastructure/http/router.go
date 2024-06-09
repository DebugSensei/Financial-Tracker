package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	handler := NewHandler(db)

	router.POST("/transaction", handler.AddTransaction)
	router.GET("/balance", handler.GetBalance)
	router.POST("/category", handler.AddCategory)
	router.GET("/categories", handler.GetCategories)
	router.GET("/transactions", handler.GetTransactions)

	return router
}
