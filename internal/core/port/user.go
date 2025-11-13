package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
}

type UserService interface {
	RegisterNewUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
}