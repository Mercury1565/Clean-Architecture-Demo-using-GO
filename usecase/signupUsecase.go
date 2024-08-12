package usecase

import (
	"Clean_Architecture/domain"
	"Clean_Architecture/utils"
	"context"
	"time"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUseCase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (signupUC *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, signupUC.contextTimeout)
	defer cancel()

	return signupUC.userRepository.Create(ctx, user)
}

func (signupUC *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, signupUC.contextTimeout)
	defer cancel()

	return signupUC.userRepository.GetByEmail(ctx, email)
}

func (signupUC *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (signupUC *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
