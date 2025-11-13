package repository

import (
	"context"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type EnrollmentRepository struct {
	db *postgres.DB
}

func InitEnrollmentRepository(db *postgres.DB) *EnrollmentRepository {
	return &EnrollmentRepository{
		db: db,
	}
}

func (er *EnrollmentRepository) CreateEnrollment(ctx context.Context, UserID, TrainingID uint, EnrolledAt time.Time) error {
	db := er.db.GetDB()

	if err := db.Create(&domain.Enrollment{
		UserID: UserID,
		TrainingID: TrainingID,
		EnrolledAt: EnrolledAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (er *EnrollmentRepository) GetEnrollments(ctx context.Context) ([]domain.Enrollment, error) {
	db := er.db.GetDB()

	var enrollments []domain.Enrollment
	if err := db.Preload("User").Preload("Training").Find(&enrollments).Error; err != nil {
		return nil, err
	}

	return enrollments, nil
}

func (er *EnrollmentRepository) GetEnrollmentsByEmail(ctx context.Context, email string) ([]domain.Enrollment, error) {
	db := er.db.GetDB()

	var enrollments []domain.Enrollment
	if err := db.Table("enrollments").Joins("left join users on users.id = enrollments.user_id").Preload("User").Preload("Training").Find(&enrollments, "users.email = ?", email).Error; err != nil {
		return nil, err
	}

	return enrollments, nil
}
