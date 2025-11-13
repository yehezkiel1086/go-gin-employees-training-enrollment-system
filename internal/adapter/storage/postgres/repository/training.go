package repository

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type TrainingRepository struct {
	db *postgres.DB
}

func InitTrainingRepository(db *postgres.DB) *TrainingRepository {
	return &TrainingRepository{
		db: db,
	}
}

func (tr *TrainingRepository) CreateTraining(ctx context.Context, training *domain.Training) (*domain.Training, error) {
	db := tr.db.GetDB()

	if err := db.Create(training).Error; err != nil {
		return nil, err
	}

	return training, nil
}

func (tr *TrainingRepository) GetTrainings(ctx context.Context) ([]domain.Training, error) {
	db := tr.db.GetDB()

	var trainings []domain.Training
	if err := db.Find(&trainings).Error; err != nil {
		return nil, err
	}

	return trainings, nil
}

func (tr *TrainingRepository) GetTrainingByID(ctx context.Context, id string) (*domain.Training, error) {
	db := tr.db.GetDB()

	var training domain.Training
	if err := db.First(&training, id).Error; err != nil {
		return nil, err
	}

	return &training, nil
}
