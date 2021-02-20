package usecase

import (
	"context"
	"reflect"
	"hub-service/domain"
	"testing"
)

func TestGetAllUserInteractor_Execute(t *testing.T) {
	t.Parallel()

	type args struct {
		input domain.FindUsersRequest
	}

	tests := []struct {
		name          string
		args          args
		repository    domain.UserRepository
		expected      []*UserOutput
		expectedError interface{}
	}{
		{
			name: "Get all user successful",
			args: args{
				input: domain.FindUsersRequest{},
			},
			repository: mockUserRepo{
				items: []*domain.User{
					{Id: 1, Email: "aaa@bbb.com"}, {Id: 2, Email: "aaa@bbb.com"}, {Id: 3},
				},
				err: nil,
			},
			expected: []*UserOutput{
				{ID: 1, Email: "aaa@bbb.com"}, {ID: 2, Email: "aaa@bbb.com"}, {ID: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewGetAllUserUseCase(tt.repository)

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
