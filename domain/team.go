package domain

import "context"

const (
	TeamTypePro        TeamType = "pro"
	TeamTypeBackOffice TeamType = "back_office"
	TypeTypeStaff      TeamType = "staff"
)

var (
	AvailableTypes = []TeamType{
		TeamTypePro,
		TeamTypeBackOffice,
		TypeTypeStaff,
	}
)

type (
	TeamType string

	FindTeamsRequest struct {
		Ids   []int64
		Name  string
		Types []string
		HubId int64
	}

	Team struct {
		Id    int64    `json:"id"`
		Name  string   `json:"name"`
		Type  TeamType `json:"type"`
		HubId int64    `json:"hub_id"`
	}

	TeamRepository interface {
		Create(context.Context, Team) (*Team, error)
		FindById(ctx context.Context, teamId int64) (*Team, error)
		Find(ctx context.Context, req FindTeamsRequest) ([]*Team, error)
		AddUsers(ctx context.Context, teamId int64, userIds []int64) error
	}
)
