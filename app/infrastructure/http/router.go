package http

import (
	"database/sql"
	"financial_tracker/app/ports"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	handler := ports.NewHandler(db)

	router.POST("/transactions", handler.AddTransaction)
	router.GET("/balance", handler.GetBalance)
	router.POST("/categories", handler.AddCategory)
	router.GET("/categories", handler.GetCategories)
	router.GET("/transactions", handler.GetTransactions)

	return router
}
