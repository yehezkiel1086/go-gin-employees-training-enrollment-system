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
	trainingByIDCachePrefix = "training:id"
	allTrainingsCacheKey = "trainings:all"
)

type TrainingService struct {
	repo       port.TrainingRepository
	cache *redis.Redis
}

func InitTrainingService(repo port.TrainingRepository, cache *redis.Redis) *TrainingService {
	return &TrainingService{
		repo:       repo,
		cache: cache,
	}
}

func (ts *TrainingService) CreateTraining(ctx context.Context, training *domain.Training) (*domain.Training, error) {
	createdTraining, err := ts.repo.CreateTraining(ctx, training)
	if err != nil {
		return nil, err
	}

	// Invalidate caches
	ts.cache.Del(ctx, allTrainingsCacheKey)

	return createdTraining, nil
}

func (ts *TrainingService) GetTrainings(ctx context.Context) ([]domain.Training, error) {
	// Try to get from cache first
	cachedTrainings, err := ts.cache.Get(ctx, allTrainingsCacheKey)
	if err == nil {
		var trainings []domain.Training
		if err := util.Deserialize(cachedTrainings, &trainings); err == nil {
			return trainings, nil
		}
	}

	// Get from repo if not in cache
	trainings, err := ts.repo.GetTrainings(ctx)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if trainingsBytes, err := util.Serialize(trainings); err == nil {
		ts.cache.Set(ctx, allTrainingsCacheKey, trainingsBytes, 1*time.Hour)
	}

	return trainings, nil
}

func (ts *TrainingService) GetTrainingByID(ctx context.Context, id string) (*domain.Training, error) {
	cacheKey := util.GenerateCacheKey(trainingByIDCachePrefix, id)

	// Try to get from cache first
	cachedTraining, err := ts.cache.Get(ctx, cacheKey)
	if err == nil {
		var training domain.Training
		if err := util.Deserialize(cachedTraining, &training); err == nil {
			return &training, nil
		}
	}

	// Get from repo if not in cache
	training, err := ts.repo.GetTrainingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if trainingBytes, err := util.Serialize(training); err == nil {
		ts.cache.Set(ctx, cacheKey, trainingBytes, 1*time.Hour)
	}

	return training, nil
}

func (ts *TrainingService) DeleteTraining(ctx context.Context, training *domain.Training) error {
	err := ts.repo.DeleteTraining(ctx, training)
	if err != nil {
		return err
	}

	// Invalidate caches
	ts.cache.Del(ctx, allTrainingsCacheKey)
	ts.cache.Del(ctx, util.GenerateCacheKey(trainingByIDCachePrefix, fmt.Sprint(training.ID)))

	return nil
}

func (ts *TrainingService) UpdateTraining(ctx context.Context, training *domain.Training) (*domain.Training, error) {
	updatedTraining, err := ts.repo.UpdateTraining(ctx, training)
	if err != nil {
		return nil, err
	}

	// Invalidate caches
	ts.cache.Del(ctx, allTrainingsCacheKey)
	ts.cache.Del(ctx, util.GenerateCacheKey(trainingByIDCachePrefix, fmt.Sprint(training.ID)))

	return updatedTraining, nil
}
