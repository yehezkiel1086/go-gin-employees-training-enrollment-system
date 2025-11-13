package repository

import (
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type StatisticsRepository struct {
	db *postgres.DB
}

func InitStatisticsRepository(db *postgres.DB) *StatisticsRepository {
	return &StatisticsRepository{
		db: db,
	}
}

func (sr *StatisticsRepository) GetTrainingStatistics() (*domain.TrainingStatistics, error) {
	db := sr.db.GetDB()
	var stats domain.TrainingStatistics

	if err := db.Model(&domain.Training{}).Count(&stats.TotalAvailableTrainings).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&domain.Enrollment{}).Count(&stats.EnrolledTrainings).Error; err != nil {
		return nil, err
	}

	// training is completed if its end date (date + duration) has passed
	if err := db.Model(&domain.Training{}).Where("date + (duration * interval '1 day') < ?", time.Now()).Count(&stats.CompletedTrainings).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}