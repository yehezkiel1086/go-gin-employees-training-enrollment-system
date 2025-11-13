package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type TrainingService struct {
	repo port.TrainingRepository
}

func InitTrainingService(repo port.TrainingRepository) *TrainingService {
	return &TrainingService{
		repo: repo,
	}
}

func (ts *TrainingService) CreateTraining(ctx context.Context, training *domain.Training) (*domain.Training, error) {
	return ts.repo.CreateTraining(ctx, training)
}

func (ts *TrainingService) GetTrainings(ctx context.Context) ([]domain.Training, error) {
	return ts.repo.GetTrainings(ctx)
}

func (ts *TrainingService) GetTrainingByID(ctx context.Context, id string) (*domain.Training, error) {
	return ts.repo.GetTrainingByID(ctx, id)
}
