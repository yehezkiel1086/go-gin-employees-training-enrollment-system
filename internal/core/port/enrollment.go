package port

import (
	"context"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type EnrollmentRepository interface {
	CreateEnrollment(ctx context.Context, UserID, TrainingID uint, EnrolledAt time.Time) error
	GetEnrollments(ctx context.Context) ([]domain.Enrollment, error)
	GetEnrollmentsByEmail(ctx context.Context, email string) ([]domain.Enrollment, error)
}

type EnrollmentService interface {
	CreateEnrollment(ctx context.Context, Email string, TrainingID uint, EnrolledAt time.Time) error
	GetEnrollments(ctx context.Context) ([]domain.Enrollment, error)
	GetEnrollmentsByEmail(ctx context.Context, email string) ([]domain.Enrollment, error)
}