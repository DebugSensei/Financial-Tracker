package services

import (
	"financial-tracker/interfaces/repository"
	"financial-tracker/internal/domain/category"
)

type CategoryService struct {
	CategoryRepo repository.CategoryRepository
}

func (s *CategoryService) AddCategory(name string) (category.Category, error) {
	cat := category.Category{Name: name}
	err := s.CategoryRepo.AddCategory(cat)
	return cat, err
}

func (s *CategoryService) GetCategories() ([]category.Category, error) {
	return s.CategoryRepo.GetCategories()
}
