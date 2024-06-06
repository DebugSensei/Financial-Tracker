package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/transaction", AddTransaction)
	router.GET("/balance", GetBalance)
	router.POST("/category", AddCategory)
	router.GET("/categories", GetCategories)
	router.GET("/transactions", GetTransactions)

	return router
}
