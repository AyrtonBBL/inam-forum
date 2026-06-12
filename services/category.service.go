package services

import (
	"inam-forum/models"
	"inam-forum/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func InitCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}
