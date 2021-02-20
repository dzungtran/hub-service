package usecase

import (
	"context"
	"errors"
	"reflect"
	"hub-service/domain"
	"testing"
)

func TestCreateUserInteractor_Execute(t *testing.T) {
	t.Parallel()

	type args struct {
		input CreateUserInput
	}

	tests := []struct {
		name          string
		args          args
		repository    domain.UserRepository
		expected      *CreateUserOutput
		expectedError interface{}
	}{
		{
			name: "Create user successful",
			args: args{
				input: CreateUserInput{
					Email: "Test@email.com",
					Role: "admin",
				},
			},
			repository: mockUserRepo{
				result: &domain.User{
					Id:   mockDefaultId,
					Email: "Test@email.com",
					Role: "admin",
				},
				err: nil,
			},
			expected: &CreateUserOutput{
				ID:   mockDefaultId,
				Email: "Test@email.com",
				Role: "admin",
			},
		},
		{
			name: "Create user successful",
			args: args{
				input: CreateUserInput{
					Email: "Test@email.com",
					Role: "admin",
				},
			},
			repository: mockUserRepo{
				result: &domain.User{
					Id:          mockDefaultId,
					Email: "Test@email.com",
					Role: "admin",
				},
				err: nil,
			},
			expected: &CreateUserOutput{
				ID:   mockDefaultId,
				Email: "Test@email.com",
				Role: "admin",
			},
		},
		{
			name: "Create user and return generic error",
			args: args{
				input: CreateUserInput{
					Email: "Test@email.com",
					Role: "admin",
				},
			},
			repository: mockUserRepo{
				result: nil,
				err: errors.New("error"),
			},
			expected: nil,
			expectedError: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateUserUseCase(tt.repository)

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
