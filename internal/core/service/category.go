package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/redis"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/util"
)

const (
	categoryByIDCachePrefix = "category:id"
	allCategoriesCacheKey = "categories:all"
)

type CategoryService struct {
	repo       port.CategoryRepository
	cache *redis.Redis
}

func InitCategoryService(repo port.CategoryRepository, cache *redis.Redis) *CategoryService {
	return &CategoryService{
		repo:       repo,
		cache: cache,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	createdCategory, err := cs.repo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	cs.cache.Del(ctx, allCategoriesCacheKey)

	return createdCategory, nil
}

func (cs *CategoryService) GetCategories(ctx context.Context) ([]domain.Category, error) {
	// Try to get from cache first
	cachedCategories, err := cs.cache.Get(ctx, allCategoriesCacheKey)
	if err == nil {
		var categories []domain.Category
		if err := util.Deserialize(cachedCategories, &categories); err == nil {
			return categories, nil
		}
	}

	// Get from repo if not in cache
	categories, err := cs.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if categoriesBytes, err := util.Serialize(categories); err == nil {
		cs.cache.Set(ctx, allCategoriesCacheKey, categoriesBytes, 1*time.Hour)
	}

	return categories, nil
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, category *domain.Category) error {
	err := cs.repo.DeleteCategory(ctx, category)
	if err != nil {
		return err
	}

	// Invalidate caches
	cs.cache.Del(ctx, allCategoriesCacheKey)
	cs.cache.Del(ctx, util.GenerateCacheKey(categoryByIDCachePrefix, fmt.Sprint(category.ID)))

	return nil
}

func (cs *CategoryService) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {
	cacheKey := util.GenerateCacheKey(categoryByIDCachePrefix, id)

	// Try to get from cache first
	cachedCategory, err := cs.cache.Get(ctx, cacheKey)
	if err == nil {
		var category domain.Category
		if err := util.Deserialize(cachedCategory, &category); err == nil {
			return &category, nil
		}
	}

	// Get from repo if not in cache
	category, err := cs.repo.GetCategoryByID(ctx, id) // This returns (*domain.Category, error)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if categoryBytes, err := util.Serialize(category); err == nil {
		cs.cache.Set(ctx, cacheKey, categoryBytes, 1*time.Hour)
	}

	return category, nil
}