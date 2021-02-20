package usecase

import (
	"context"
	"errors"
	"reflect"
	"hub-service/domain"
	"testing"
)

func TestCreateHubInteractor_Execute(t *testing.T) {
	t.Parallel()

	type args struct {
		input CreateHubInput
	}

	tests := []struct {
		name          string
		args          args
		repository    domain.HubRepository
		expected      *CreateHubOutput
		expectedError interface{}
	}{
		{
			name: "Create hub successful",
			args: args{
				input: CreateHubInput{
					Name: "Test",
					Lat:  88,
					Long: 88,
				},
			},
			repository: mockHubRepo{
				result: &domain.Hub{
					Id:   mockDefaultId,
					Name: "Test",
					GeoLocation: domain.GeoLocation{
						Lat:  88,
						Long: 88,
					},
				},
				err: nil,
			},
			expected: &CreateHubOutput{
				ID:   mockDefaultId,
				Name: "Test",
				Lat:  88,
				Long: 88,
			},
		},
		{
			name: "Create hub successful",
			args: args{
				input: CreateHubInput{
					Name: "Test",
				},
			},
			repository: mockHubRepo{
				result: &domain.Hub{
					Id:          mockDefaultId,
					Name:        "Test",
					GeoLocation: domain.GeoLocation{},
				},
				err: nil,
			},
			expected: &CreateHubOutput{
				ID:   mockDefaultId,
				Name: "Test",
				Lat:  0,
				Long: 0,
			},
		},
		{
			name: "Create hub and return generic error",
			args: args{
				input: CreateHubInput{
					Name: "Test",
				},
			},
			repository: mockHubRepo{
				result: nil,
				err: errors.New("error"),
			},
			expected: nil,
			expectedError: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateHubUseCase(tt.repository)

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
