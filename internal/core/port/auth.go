package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, userReq *domain.User) (*domain.User, error)
}
