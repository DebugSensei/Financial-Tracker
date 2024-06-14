package http

import (
	"database/sql"
	"financial_tracker/app/ports"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) AddTransaction(c *gin.Context) {
	ports.AddTransaction(h.DB, c)
}

func (h *Handler) GetBalance(c *gin.Context) {
	ports.GetBalance(h.DB, c)
}

func (h *Handler) AddCategory(c *gin.Context) {
	ports.AddCategory(h.DB, c)
}

func (h *Handler) GetCategories(c *gin.Context) {
	ports.GetCategories(h.DB, c)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	ports.GetTransactions(h.DB, c)
}
