package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/util"
)

type AuthService struct {
	userRepo port.UserRepository
}

func InitAuthService(userRepo port.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, userReq *domain.User) (*domain.User, error) {
	// check email
	user, err := as.userRepo.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		return nil, err
	}

	// check password
	if err := util.ComparePassword(user.Password, userReq.Password); err != nil {
		return nil, err
	}

	return user, nil
}
