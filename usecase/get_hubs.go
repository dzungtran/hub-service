package usecase

import (
	"context"
	"hub-service/domain"
)

type (
	GetHubsUseCase interface {
		Execute(context.Context, GetHubsInput) ([]*HubOutput, error)
	}

	GetHubsInput struct {
		Ids  []int64
		Name string
	}

	// Output data
	HubOutput struct {
		ID   int64   `json:"id"`
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}

	getAllHubInteractor struct {
		hubRepo domain.HubRepository
	}
)

func NewGetHubsUseCase(hubRepo domain.HubRepository) GetHubsUseCase {
	return getAllHubInteractor{
		hubRepo: hubRepo,
	}
}

// Execute create Hub with dependencies
func (i getAllHubInteractor) Execute(ctx context.Context, req GetHubsInput) ([]*HubOutput, error) {

	teams, err := i.hubRepo.Find(ctx, domain.FindHubsRequest{
		Ids:  req.Ids,
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return transformHubsToSliceHubOutput(teams), nil
}

func transformHubsToSliceHubOutput(hubs []*domain.Hub) []*HubOutput {
	result := make([]*HubOutput, 0)
	for _, u := range hubs {
		result = append(result, &HubOutput{
			ID:   u.Id,
			Name: u.Name,
			Lat:  u.GeoLocation.Lat,
			Long: u.GeoLocation.Long,
		})
	}
	return result
}
