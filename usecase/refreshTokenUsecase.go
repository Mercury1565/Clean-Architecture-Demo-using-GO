package usecase

import (
	"Clean_Architecture/domain"
	"Clean_Architecture/utils"
	"context"
	"time"
)

type refreshTokenUsecase struct {
	userRepostory  domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepostory:  userRepository,
		contextTimeout: timeout,
	}
}

func (refreshTokenUC *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, refreshTokenUC.contextTimeout)
	defer cancel()

	return refreshTokenUC.userRepostory.GetByID(ctx, id)
}

func (refreshTokenUC *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return utils.ExtractIDFromToken(requestToken, secret)
}

func (refreshTokenUC *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (refreshTokenUC *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
