package repository

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type CategoryRepository struct {
	db *postgres.DB
}

func InitCategoryRepository(db *postgres.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	db := cr.db.GetDB()
	if err := db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *CategoryRepository) GetCategories(ctx context.Context) ([]domain.Category, error) {
	db := cr.db.GetDB()

	var categories []domain.Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepository) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {
	db := cr.db.GetDB()

	var category domain.Category
	if err := db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (cr *CategoryRepository) DeleteCategory(ctx context.Context, category *domain.Category) error {
	db := cr.db.GetDB()
	if err := db.Delete(category).Error; err != nil {
		return err
	}

	return nil
}
