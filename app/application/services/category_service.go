package services

import (
	"financial_tracker/app/interfaces/repository"
	"financial_tracker/app/internal/domain/category"
)

// CategoryService struct
type CategoryService struct {
	CategoryRepo repository.CategoryRepository
}

// NewCategoryService creates a new CategoryService
func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: repo}
}

func (s *CategoryService) AddCategory(cat category.Category) error {
	return s.CategoryRepo.AddCategory(cat)
}

func (s *CategoryService) GetCategories() ([]category.Category, error) {
	return s.CategoryRepo.GetCategories()
}
