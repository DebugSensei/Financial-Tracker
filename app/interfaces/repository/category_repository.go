package repository

import (
	"database/sql"
	"financial_tracker/infrastructure/db"
	"financial_tracker/internal/domain/category"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: db.DB}
}

func (r *CategoryRepository) AddCategory(category category.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1)`
	_, err := r.db.Exec(query, category.Name)
	return err
}

func (r *CategoryRepository) GetCategories() ([]category.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []category.Category
	for rows.Next() {
		var category category.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
