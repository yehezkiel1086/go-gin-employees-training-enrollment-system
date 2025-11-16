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
	userByEmailCachePrefix = "user:email"
	allUsersCacheKey       = "users:all"
)

type UserService struct {
	repo       port.UserRepository
	cache *redis.Redis
}

func InitUserService(repo port.UserRepository, cache *redis.Redis) *UserService {
	return &UserService{
		repo: repo,
		cache: cache,
	}
}

func (us *UserService) RegisterNewUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPwd

	createdUser, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// Invalidate cache for all users
	us.cache.Del(ctx, allUsersCacheKey)

	return createdUser, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	cacheKey := util.GenerateCacheKey(userByEmailCachePrefix, email)

	// Try to get from cache first
	cachedUser, err := us.cache.Get(ctx, cacheKey)
	if err == nil {
		var user domain.User
		if err := util.Deserialize(cachedUser, &user); err == nil {
			return &user, nil
		}
	}

	// Get from repo if not in cache
	user, err := us.repo.GetUserByEmail(ctx, email) // This returns (*domain.User, error)
	if err != nil {
		return nil, err
	}

	// Set to cache
	if userBytes, err := util.Serialize(user); err == nil {
		us.cache.Set(ctx, cacheKey, userBytes, 1*time.Hour)
	}

	return user, nil
}

func (us *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	cachedUsers, err := us.cache.Get(ctx, allUsersCacheKey)
	if err == nil {
		var users []domain.User
		if err := util.Deserialize(cachedUsers, &users); err == nil {
			return users, nil
		}
	}

	users, err := us.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	if usersBytes, err := util.Serialize(users); err == nil {
		us.cache.Set(ctx, allUsersCacheKey, usersBytes, 1*time.Hour)
	}

	return users, nil
}
