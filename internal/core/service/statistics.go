package service

import (
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type StatisticsService struct {
	repo port.StatisticsRepository
}

func InitStatisticsService(repo port.StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		repo: repo,
	}
}

func (ss *StatisticsService) GetTrainingStatistics() (*domain.TrainingStatistics, error) {
	return ss.repo.GetTrainingStatistics()
}

func (ss *StatisticsService) GetTrainingsByCategories() ([]domain.TrainingsByCategory, error) {
	return ss.repo.GetTrainingsByCategories()
}
