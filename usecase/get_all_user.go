package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	GetAllUserUseCase interface {
		Execute(ctx context.Context, req domain.FindUsersRequest) ([]*UserOutput, error)
	}

	// Output data
	UserOutput struct {
		ID    int64  `json:"id"`
		Role  string `json:"role"`
		Email string `json:"email"`
	}

	getAllUserInteractor struct {
		userRepo domain.UserRepository
	}
)

func NewGetAllUserUseCase(userRepo domain.UserRepository) GetAllUserUseCase {
	return getAllUserInteractor{
		userRepo: userRepo,
	}
}

// Execute create User with dependencies
func (i getAllUserInteractor) Execute(ctx context.Context, req domain.FindUsersRequest) ([]*UserOutput, error) {

	users, err := i.userRepo.Find(ctx, req)
	if err != nil {
		return nil, err
	}

	return transformUsersToSliceUserOutput(users), nil
}

func transformUsersToSliceUserOutput(users []*domain.User) []*UserOutput {
	result := make([]*UserOutput, 0)
	for _, u := range users {
		result = append(result, &UserOutput{
			ID:    u.Id,
			Role:  string(u.Role),
			Email: u.Email,
		})
	}
	return result
}
