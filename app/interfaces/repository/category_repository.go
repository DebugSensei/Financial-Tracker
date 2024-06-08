package repository

import (
	"database/sql"
	"financial_tracker/app/internal/domain/category"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) AddCategory(cat category.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1)`
	_, err := r.db.Exec(query, cat.Name)
	return err
}

func (r *CategoryRepository) GetCategories() ([]category.Category, error) {
	var categories []category.Category
	query := `SELECT id, name FROM categories`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat category.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}
