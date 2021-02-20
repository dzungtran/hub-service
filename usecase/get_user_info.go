package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	GetUserInfoUseCase interface {
		Execute(ctx context.Context, userId int64) (*UserOutput, error)
	}

	getUserInfoInteractor struct {
		userRepo domain.UserRepository
	}
)

func NewGetUserInfoUseCase(userRepo domain.UserRepository) GetUserInfoUseCase {
	return getUserInfoInteractor{
		userRepo: userRepo,
	}
}

// Execute create User with dependencies
func (i getUserInfoInteractor) Execute(ctx context.Context, userId int64) (*UserOutput, error) {

	user, err := i.userRepo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return transformUsersToSliceUserOutput([]*domain.User{user})[0], nil
}
