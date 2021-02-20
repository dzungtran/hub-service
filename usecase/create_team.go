package usecase

import (
	"context"
	"errors"
	"hub-service/domain"
)

type (
	CreateTeamUseCase interface {
		Execute(context.Context, CreateTeamInput) (*CreateTeamOutput, error)
	}

	// Input data
	CreateTeamInput struct {
		Name  string `json:"name" validate:"required"`
		Type  string `json:"type" validate:"required"`
		HubId int64  `json:"hub_id" validate:"required"`
	}

	// Output data
	CreateTeamOutput struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		HubId int64  `json:"hub_id"`
	}

	createTeamInteractor struct {
		teamRepo domain.TeamRepository
		hubRepo  domain.HubRepository
	}
)

func NewCreateTeamUseCase(teamRepo domain.TeamRepository, hubRepo domain.HubRepository) CreateTeamUseCase {
	return createTeamInteractor{
		teamRepo: teamRepo,
		hubRepo:  hubRepo,
	}
}

// Execute create Team with dependencies
func (i createTeamInteractor) Execute(ctx context.Context, input CreateTeamInput) (*CreateTeamOutput, error) {
	var err error

	// check hub_id is exists
	_, err = i.hubRepo.FindById(ctx, input.HubId)
	if err != nil {
		if err == domain.ErrorNotFound {
			return nil, errors.New("Hub not found")
		}
		return nil, err
	}

	team, err := i.teamRepo.Create(ctx, transformCreateTeamInputToTeamObject(input))
	if err != nil {
		return nil, err
	}

	return transformTeamObjectToCreateTeamOutput(team), nil
}

func transformCreateTeamInputToTeamObject(input CreateTeamInput) domain.Team {
	return domain.Team{
		Name:  input.Name,
		Type:  domain.TeamType(input.Type),
		HubId: input.HubId,
	}
}

func transformTeamObjectToCreateTeamOutput(team *domain.Team) *CreateTeamOutput {
	return &CreateTeamOutput{
		Id:    team.Id,
		Name:  team.Name,
		Type:  string(team.Type),
		HubId: team.HubId,
	}
}
