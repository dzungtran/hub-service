package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (*CreateUserOutput, error)
	}

	// Input data
	CreateUserInput struct {
		Role  string `json:"role" validate:"required"`
		Email string `json:"email" validate:"required"`
	}

	// Output data
	CreateUserOutput struct {
		ID    int64  `json:"id"`
		Role  string `json:"role"`
		Email string `json:"email"`
	}

	createUserInteractor struct {
		repo domain.UserRepository
	}
)

func NewCreateUserUseCase(repo domain.UserRepository) CreateUserUseCase {
	return createUserInteractor{
		repo: repo,
	}
}

// Execute create User with dependencies
func (i createUserInteractor) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	var err error

	user, err := i.repo.Create(ctx, transformCreateUserInputToUserObject(input))
	if err != nil {
		return nil, err
	}

	return transformUserObjectToCreateUserOutput(user), nil
}

func transformCreateUserInputToUserObject(input CreateUserInput) domain.User {
	return domain.User{
		Role:  domain.UserRoleType(input.Role),
		Email: input.Email,
	}
}

func transformUserObjectToCreateUserOutput(user *domain.User) *CreateUserOutput {
	return &CreateUserOutput{
		ID:    user.Id,
		Role:  string(user.Role),
		Email: user.Email,
	}
}
