package usecase

import (
	"Clean_Architecture/domain"
	"Clean_Architecture/utils"
	"context"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUseCase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (loginUC *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, loginUC.contextTimeout)
	defer cancel()

	return loginUC.userRepository.GetByEmail(ctx, email)
}

func (loginUsecase *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (loginUsecase *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
