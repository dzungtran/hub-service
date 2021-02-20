package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	GetHubInfoUseCase interface {
		Execute(ctx context.Context, teamId int64) (*HubOutput, error)
	}

	getHubInfoInteractor struct {
		teamRepo domain.HubRepository
	}
)

func NewGetHubInfoUseCase(teamRepo domain.HubRepository) GetHubInfoUseCase {
	return getHubInfoInteractor{
		teamRepo: teamRepo,
	}
}

// Execute create Hub with dependencies
func (i getHubInfoInteractor) Execute(ctx context.Context, teamId int64) (*HubOutput, error) {

	team, err := i.teamRepo.FindById(ctx, teamId)
	if err != nil {
		return nil, err
	}

	return transformHubsToSliceHubOutput([]*domain.Hub{team})[0], nil
}
