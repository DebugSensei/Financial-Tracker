// models/category.go
package models

// Category represents a category of transactions
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
