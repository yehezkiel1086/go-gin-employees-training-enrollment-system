package service

import (
	"context"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type EnrollmentService struct {
	repo port.EnrollmentRepository
	userRepo port.UserRepository
}

func InitEnrollmentService(repo port.EnrollmentRepository, userRepo port.UserRepository) *EnrollmentService {
	return &EnrollmentService{
		repo: repo,
		userRepo: userRepo,
	}
}

func (es *EnrollmentService) CreateEnrollment(ctx context.Context, Email string, TrainingID uint, EnrolledAt time.Time) error {
	// get user to get the id
	user, err := es.userRepo.GetUserByEmail(ctx, Email)
	if err != nil {
		return err
	}

	return es.repo.CreateEnrollment(ctx, user.ID, TrainingID, EnrolledAt)
}

func (es *EnrollmentService) GetEnrollments(ctx context.Context) ([]domain.Enrollment, error) {
	return es.repo.GetEnrollments(ctx)
}

func (es *EnrollmentService) GetEnrollmentsByEmail(ctx context.Context, email string) ([]domain.Enrollment, error) {
	return es.repo.GetEnrollmentsByEmail(ctx, email)
}
