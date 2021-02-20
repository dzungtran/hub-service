package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	GetTeamInfoUseCase interface {
		Execute(ctx context.Context, teamId int64) (*TeamOutput, error)
	}

	getTeamInfoInteractor struct {
		teamRepo domain.TeamRepository
	}
)

func NewGetTeamInfoUseCase(teamRepo domain.TeamRepository) GetTeamInfoUseCase {
	return getTeamInfoInteractor{
		teamRepo: teamRepo,
	}
}

// Execute create Team with dependencies
func (i getTeamInfoInteractor) Execute(ctx context.Context, teamId int64) (*TeamOutput, error) {

	team, err := i.teamRepo.FindById(ctx, teamId)
	if err != nil {
		return nil, err
	}

	return transformTeamsToSliceTeamOutput([]*domain.Team{team})[0], nil
}
