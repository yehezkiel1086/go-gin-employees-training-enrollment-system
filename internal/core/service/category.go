package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type CategoryService struct {
	repo port.CategoryRepository
}

func InitCategoryService(repo port.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return cs.repo.CreateCategory(ctx, category)
}

func (cs *CategoryService) GetCategories(ctx context.Context) ([]domain.Category, error) {
	return cs.repo.GetCategories(ctx)
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, category *domain.Category) error {
	return cs.repo.DeleteCategory(ctx, category)
}

func (cs *CategoryService) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {
	return cs.repo.GetCategoryByID(ctx, id)
}