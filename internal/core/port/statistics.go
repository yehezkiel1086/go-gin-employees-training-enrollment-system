package port

import "github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"

type StatisticsRepository interface {
	GetTrainingStatistics() (*domain.TrainingStatistics, error)
}

type StatisticsService interface {
	GetTrainingStatistics() (*domain.TrainingStatistics, error)
}