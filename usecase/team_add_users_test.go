package usecase

import (
	"context"
	"errors"
	"hub-service/domain"
	"testing"
)

func TestTeamAddUserInteractor_Execute(t *testing.T) {
	t.Parallel()

	type args struct {
		input TeamAddUsersInput
	}

	tests := []struct {
		name           string
		args           args
		teamRepository domain.TeamRepository
		userRepository domain.UserRepository
		expectedError  interface{}
	}{
		{
			name: "Team add users successful",
			args: args{
				input: TeamAddUsersInput{
					UserIds: []int64{1, 2, 3},
					TeamId:  mockDefaultId,
				},
			},
			teamRepository: mockTeamRepo{
				result: &domain.Team{
					Id:    mockDefaultId,
					Name:  "Test",
					HubId: mockDefaultId,
				},
				err: nil,
			},
			userRepository: mockUserRepo{
				items: []*domain.User{
					{Id: 1}, {Id: 2}, {Id: 3},
				},
				err: nil,
			},
		},
		{
			name: "Team add users with some users does not exists then return error",
			args: args{
				input: TeamAddUsersInput{
					UserIds: []int64{1, 2, 3},
					TeamId:  mockDefaultId,
				},
			},
			teamRepository: mockTeamRepo{
				result: &domain.Team{
					Id:    mockDefaultId,
					Name:  "Test",
					HubId: mockDefaultId,
				},
				err: nil,
			},
			userRepository: mockUserRepo{
				items: []*domain.User{
					{Id: 1}, {Id: 2},
				},
				err: nil,
			},
			expectedError: "User Ids [3] does not exists",
		},
		{
			name: "Team add users with error while fetch users then return error",
			args: args{
				input: TeamAddUsersInput{
					UserIds: []int64{1, 2, 3},
					TeamId:  mockDefaultId,
				},
			},
			teamRepository: mockTeamRepo{
				result: &domain.Team{
					Id:    mockDefaultId,
					Name:  "Test",
					HubId: mockDefaultId,
				},
				err: nil,
			},
			userRepository: mockUserRepo{
				items: nil,
				err:   errors.New("error"),
			},
			expectedError: "error",
		},
		{
			name: "Team add users then return generic error",
			args: args{
				input: TeamAddUsersInput{
					UserIds: []int64{1, 2, 3},
					TeamId:  mockDefaultId,
				},
			},
			teamRepository: mockTeamRepo{
				result: nil,
				err:    errors.New("error"),
			},
			userRepository: mockUserRepo{
				items: []*domain.User{
					{Id: 1}, {Id: 2}, {Id: 3},
				},
				err: nil,
			},
			expectedError: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewTeamAddUsersUseCase(tt.teamRepository, tt.userRepository)

			err := uc.Execute(context.TODO(), tt.args.input)
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.name, err, tt.expectedError)
			}
		})
	}
}
