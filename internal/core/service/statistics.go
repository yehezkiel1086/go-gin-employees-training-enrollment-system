package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/redis"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

const (
	trainingStatsCacheKey        = "stats:trainings"
	trainingsByCatStatsCacheKey = "stats:trainings-by-categories"
)

type StatisticsService struct {
	repo       port.StatisticsRepository
	cache *redis.Redis
}

func InitStatisticsService(repo port.StatisticsRepository, cache *redis.Redis) *StatisticsService {
	return &StatisticsService{
		repo:       repo,
		cache: cache,
	}
}

func (ss *StatisticsService) GetTrainingStatistics(ctx context.Context) (*domain.TrainingStatistics, error) {
	cachedStats, err := ss.cache.Get(ctx, trainingStatsCacheKey)
	if err == nil {
		var stats domain.TrainingStatistics
		if json.Unmarshal(cachedStats, &stats) == nil {
			return &stats, nil
		}
	}

	stats, err := ss.repo.GetTrainingStatistics(ctx)
	if err != nil {
		return nil, err
	}

	statsJSON, _ := json.Marshal(stats)
	ss.cache.Set(ctx, trainingStatsCacheKey, statsJSON, 1*time.Hour)

	return stats, nil
}

func (ss *StatisticsService) GetTrainingsByCategories(ctx context.Context) ([]domain.TrainingsByCategory, error) {
	cachedStats, err := ss.cache.Get(ctx, trainingsByCatStatsCacheKey)
	if err == nil {
		var stats []domain.TrainingsByCategory
		if json.Unmarshal(cachedStats, &stats) == nil {
			return stats, nil
		}
	}
	stats, err := ss.repo.GetTrainingsByCategories(ctx)
	if err != nil {
		return nil, err
	}
	statsJSON, _ := json.Marshal(stats)
	ss.cache.Set(ctx, trainingsByCatStatsCacheKey, statsJSON, 1*time.Hour)
	return stats, nil
}
