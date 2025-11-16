package service

import (
	"context"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/redis"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/util"
)

const (
	enrollmentsByEmailCachePrefix = "enrollments:email"
	allEnrollmentsCacheKey     = "enrollments:all"
)

type EnrollmentService struct {
	repo       port.EnrollmentRepository
	userRepo   port.UserRepository
	cache *redis.Redis
}

func InitEnrollmentService(repo port.EnrollmentRepository, userRepo port.UserRepository, cache *redis.Redis) *EnrollmentService {
	return &EnrollmentService{
		repo:       repo,
		userRepo:   userRepo,
		cache: cache,
	}
}

func (es *EnrollmentService) CreateEnrollment(ctx context.Context, Email string, TrainingID uint, EnrolledAt time.Time) error {
	// get user to get the id
	user, err := es.userRepo.GetUserByEmail(ctx, Email)
	if err != nil {
		return err
	}

	err = es.repo.CreateEnrollment(ctx, user.ID, TrainingID, EnrolledAt)
	if err != nil {
		return err
	}

	// Invalidate caches
	es.cache.Del(ctx, allEnrollmentsCacheKey)
	es.cache.Del(ctx, util.GenerateCacheKey(enrollmentsByEmailCachePrefix, Email))

	return nil
}

func (es *EnrollmentService) GetEnrollments(ctx context.Context) ([]domain.Enrollment, error) {
	// Try to get from cache first
	cachedEnrollments, err := es.cache.Get(ctx, allEnrollmentsCacheKey)
	if err == nil {
		var enrollments []domain.Enrollment
		if err := util.Deserialize(cachedEnrollments, &enrollments); err == nil {
			return enrollments, nil
		}
	}

	// Get from repo if not in cache
	enrollments, err := es.repo.GetEnrollments(ctx)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if enrollmentsBytes, err := util.Serialize(enrollments); err == nil {
		es.cache.Set(ctx, allEnrollmentsCacheKey, enrollmentsBytes, 1*time.Hour)
	}

	return enrollments, nil
}

func (es *EnrollmentService) GetEnrollmentsByEmail(ctx context.Context, email string) ([]domain.Enrollment, error) {
	cacheKey := util.GenerateCacheKey(enrollmentsByEmailCachePrefix, email)
	cachedEnrollments, err := es.cache.Get(ctx, cacheKey)
	if err == nil {
		var enrollments []domain.Enrollment
		if util.Deserialize(cachedEnrollments, &enrollments) == nil {
			return enrollments, nil
		}
	}

	enrollments, err := es.repo.GetEnrollmentsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if enrollmentsBytes, err := util.Serialize(enrollments); err == nil {
		es.cache.Set(ctx, cacheKey, enrollmentsBytes, 1*time.Hour)
	}
	return enrollments, nil
}
