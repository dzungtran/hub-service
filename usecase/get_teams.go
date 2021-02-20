package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	GetTeamsUseCase interface {
		Execute(context.Context, GetTeamsInput) ([]*TeamOutput, error)
	}

	GetTeamsInput struct {
		Ids   []int64
		Name  string
		Types []string
		HubId int64
	}

	// Output data
	TeamOutput struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		HubId int64  `json:"hub_id"`
	}

	getAllTeamInteractor struct {
		teamRepo domain.TeamRepository
	}
)

func NewGetTeamsUseCase(teamRepo domain.TeamRepository) GetTeamsUseCase {
	return getAllTeamInteractor{
		teamRepo: teamRepo,
	}
}

// Execute create Team with dependencies
func (i getAllTeamInteractor) Execute(ctx context.Context, req GetTeamsInput) ([]*TeamOutput, error) {

	teams, err := i.teamRepo.Find(ctx, domain.FindTeamsRequest{
		Ids:   req.Ids,
		Types: req.Types,
		Name:  req.Name,
		HubId: req.HubId,
	})

	if err != nil {
		return nil, err
	}

	return transformTeamsToSliceTeamOutput(teams), nil
}

func transformTeamsToSliceTeamOutput(teams []*domain.Team) []*TeamOutput {
	result := make([]*TeamOutput, 0)
	for _, u := range teams {
		result = append(result, &TeamOutput{
			ID:    u.Id,
			Name:  u.Name,
			Type:  string(u.Type),
			HubId: u.HubId,
		})
	}
	return result
}
