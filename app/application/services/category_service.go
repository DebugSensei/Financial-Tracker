package services

import (
	"financial_tracker/app/internal/domain/category"
)

type CategoryRepository interface {
	AddCategory(cat category.Category) error
	GetCategories() ([]category.Category, error)
}

type CategoryService struct {
	CategoryRepo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: repo}
}

func (s *CategoryService) AddCategory(name string) (category.Category, error) {
	cat := category.Category{Name: name}
	err := s.CategoryRepo.AddCategory(cat)
	return cat, err
}

func (s *CategoryService) GetCategories() ([]category.Category, error) {
	return s.CategoryRepo.GetCategories()
}
