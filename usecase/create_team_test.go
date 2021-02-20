package usecase

import (
	"context"
	"errors"
	"reflect"
	"hub-service/domain"
	"testing"
)

func TestCreateTeamInteractor_Execute(t *testing.T) {
	t.Parallel()

	type args struct {
		input CreateTeamInput
	}

	tests := []struct {
		name           string
		args           args
		teamRepository domain.TeamRepository
		hubRepository  domain.HubRepository
		expected       *CreateTeamOutput
		expectedError  interface{}
	}{
		{
			name: "Create team successful",
			args: args{
				input: CreateTeamInput{
					Name: "Test",
					HubId: mockDefaultId,
				},
			},
			teamRepository: mockTeamRepo{
				result: &domain.Team{
					Id:   mockDefaultId,
					Name: "Test",
					HubId: mockDefaultId,
				},
				err: nil,
			},
			hubRepository: mockHubRepo{
				result: &domain.Hub{
					Id:   mockDefaultId,
				},
				err: nil,
			},
			expected: &CreateTeamOutput{
				Id:   mockDefaultId,
				Name: "Test",
				HubId: mockDefaultId,
			},
		},
		{
			name: "Create team with hub not found then return error",
			args: args{
				input: CreateTeamInput{
					Name: "Test",
				},
			},
			teamRepository: mockTeamRepo{
				result: &domain.Team{
					Id:   mockDefaultId,
					Name: "Test",
				},
				err: nil,
			},
			hubRepository: mockHubRepo{
				result: nil,
				err: domain.ErrorNotFound,
			},
			expected: nil,
			expectedError: "Hub not found",
		},
		{
			name: "Create team and return generic error",
			args: args{
				input: CreateTeamInput{
					Name: "Test",
					HubId: mockDefaultId,
				},
			},
			teamRepository: mockTeamRepo{
				result: nil,
				err: errors.New("error"),
			},
			hubRepository: mockHubRepo{
				result: &domain.Hub{
					Id:   mockDefaultId,
				},
				err: nil,
			},
			expected: nil,
			expectedError: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateTeamUseCase(tt.teamRepository, tt.hubRepository)

			result, err := uc.Execute(context.TODO(), tt.args.input)
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.name, err, tt.expectedError)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, result, tt.expected)
			}
		})
	}
}
