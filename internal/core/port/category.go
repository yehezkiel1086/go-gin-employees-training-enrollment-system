package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*domain.Category, error)
	DeleteCategory(ctx context.Context, category *domain.Category) error
}

type CategoryService interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, category *domain.	Category) error		
	GetCategoryByID(ctx context.Context, id string) (*domain.Category, error)
}