package domain

import "context"

const (
	RoleHubAdmin  UserRoleType = "hub_admin"
	RoleTeamAdmin UserRoleType = "team_admin"
	RoleCustomer  UserRoleType = "customer"
)

type (
	UserRoleType string

	User struct {
		Id    int64        `json:"id"`
		Role  UserRoleType `json:"role"`
		Email string       `json:"email"`
	}

	FindUsersRequest struct {
		Ids    []int64
		TeamId int64
	}

	UserRepository interface {
		Create(context.Context, User) (*User, error)
		FindById(ctx context.Context, userId int64) (*User, error)
		Find(ctx context.Context, req FindUsersRequest) ([]*User, error)
	}
)

var (
	AvailableRoles = []UserRoleType{
		RoleHubAdmin,
		RoleTeamAdmin,
		RoleCustomer,
	}
)
