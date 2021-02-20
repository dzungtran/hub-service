package domain

import "context"

type (
	GeoLocation struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}

	Hub struct {
		Id          int64       `json:"id"`
		Name        string      `json:"name"`
		GeoLocation GeoLocation `json:"geo_location"`
	}

	FindHubsRequest struct {
		Ids   []int64
		Name  string
	}

	HubRepository interface {
		Create(context.Context, Hub) (*Hub, error)
		FindById(ctx context.Context, hubId int64) (*Hub, error)
		Find(ctx context.Context, req FindHubsRequest) ([]*Hub, error)
	}
)
