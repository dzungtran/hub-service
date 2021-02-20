package usecase

import (
	"context"
	"hub-service/domain"
)

const mockDefaultId = int64(1000)

// Mock for HubRepository
type mockHubRepo struct {
	result *domain.Hub
	err    error
	items  []*domain.Hub
}

func (m mockHubRepo) Create(_ context.Context, _ domain.Hub) (*domain.Hub, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.result, m.err
}

func (m mockHubRepo) FindById(_ context.Context, _ int64) (*domain.Hub, error) {
	return m.result, m.err
}

func (m mockHubRepo) Find(_ context.Context, _ domain.FindHubsRequest) ([]*domain.Hub, error) {
	return m.items, m.err
}

// Mock for UserRepository
type mockUserRepo struct {
	result *domain.User
	err    error
	items  []*domain.User
}

func (m mockUserRepo) Create(_ context.Context, _ domain.User) (*domain.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.result, m.err
}

func (m mockUserRepo) FindById(_ context.Context, _ int64) (*domain.User, error) {
	return m.result, m.err
}

func (m mockUserRepo) Find(_ context.Context, _ domain.FindUsersRequest) ([]*domain.User, error) {
	return m.items, m.err
}

// Mock for TeamRepository
type mockTeamRepo struct {
	result *domain.Team
	err    error
	items  []*domain.Team
}

func (m mockTeamRepo) Create(_ context.Context, _ domain.Team) (*domain.Team, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.result, m.err
}

func (m mockTeamRepo) FindById(_ context.Context, _ int64) (*domain.Team, error) {
	return m.result, m.err
}

func (m mockTeamRepo) Find(_ context.Context, _ domain.FindTeamsRequest) ([]*domain.Team, error) {
	return m.items, m.err
}

func (m mockTeamRepo) AddUsers(_ context.Context, _ int64, _ []int64) error {
	return m.err
}
