package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type TrainingRepository interface {
	CreateTraining(ctx context.Context, training *domain.Training) (*domain.Training, error)
	GetTrainings(ctx context.Context) ([]domain.Training, error)
	GetTrainingByID(ctx context.Context, id string) (*domain.Training, error)
}

type TrainingService interface {
	CreateTraining(ctx context.Context, training *domain.Training) (*domain.Training, error)
	GetTrainings(ctx context.Context) ([]domain.Training, error)
	GetTrainingByID(ctx context.Context, id string) (*domain.Training, error)	
}