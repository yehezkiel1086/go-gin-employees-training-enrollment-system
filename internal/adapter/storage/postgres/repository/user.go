package repository

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type UserRepository struct {
	db *postgres.DB
}

func InitUserRepository(db *postgres.DB) (*UserRepository) {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := ur.db.GetDB()
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	db := ur.db.GetDB()

	var user domain.User
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	db := ur.db.GetDB()

	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
