package usecase

import (
	"Clean_Architecture/domain"
	"context"
	"time"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (profileUC *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, profileUC.contextTimeout)
	defer cancel()

	user, err := profileUC.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile := &domain.Profile{}
	profile.Name = user.Name
	profile.Email = user.Email

	return profile, nil
}
