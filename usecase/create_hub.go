package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	CreateHubUseCase interface {
		Execute(context.Context, CreateHubInput) (*CreateHubOutput, error)
	}

	// Input data
	CreateHubInput struct {
		Name string  `json:"name" validate:"required"`
		Lat  float64 `json:"lat" validate:"required"`
		Long float64 `json:"long" validate:"required"`
	}

	// Output data
	CreateHubOutput struct {
		ID   int64   `json:"id"`
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}

	createHubInteractor struct {
		repo domain.HubRepository
	}
)

func NewCreateHubUseCase(repo domain.HubRepository) CreateHubUseCase {
	return createHubInteractor{
		repo: repo,
	}
}

// Execute create User with dependencies
func (i createHubInteractor) Execute(ctx context.Context, input CreateHubInput) (*CreateHubOutput, error) {
	var err error

	hub, err := i.repo.Create(ctx, transformCreateHubInputToHubObject(input))
	if err != nil {
		return nil, err
	}

	return transformHubObjectToCreateHubOutput(hub), nil
}

func transformCreateHubInputToHubObject(input CreateHubInput) domain.Hub {
	return domain.Hub{
		Name: input.Name,
		GeoLocation: domain.GeoLocation{
			Lat:  input.Lat,
			Long: input.Long,
		},
	}
}

func transformHubObjectToCreateHubOutput(hub *domain.Hub) *CreateHubOutput {
	return &CreateHubOutput{
		ID:   hub.Id,
		Name: hub.Name,
		Lat:  hub.GeoLocation.Lat,
		Long: hub.GeoLocation.Long,
	}
}
