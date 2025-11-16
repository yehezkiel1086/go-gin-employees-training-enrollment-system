package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type StatisticsRepository interface {
	GetTrainingStatistics(ctx context.Context) (*domain.TrainingStatistics, error)
	GetTrainingsByCategories(ctx context.Context) ([]domain.TrainingsByCategory, error)
}

type StatisticsService interface {
	GetTrainingStatistics(ctx context.Context) (*domain.TrainingStatistics, error)
	GetTrainingsByCategories(ctx context.Context) ([]domain.TrainingsByCategory, error)
}