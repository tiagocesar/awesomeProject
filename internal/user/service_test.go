//nolint:lll // In tests we ignore the long-line lint

package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tiagocesar/awesomeProject/internal/repo"
)

type mockUserStorer struct{}

func (m *mockUserStorer) AddUser(u repo.User) string {
	return fmt.Sprintf("TEST - The user %s with email %s was successfully created", u.Name, u.Email)
}

func TestAddUser(t *testing.T) {
	params := map[string]string{
		"name":  "Test User",
		"email": "test@email.com",
	}

	mockUserStorer := mockUserStorer{}

	tests := []struct {
		name             string
		userStorer       userStorer
		params           map[string]string
		expectedResponse string
	}{
		{
			name:             "Success - gets response with specified name and email",
			userStorer:       &mockUserStorer,
			params:           params,
			expectedResponse: fmt.Sprintf("TEST - The user %s with email %s was successfully created", params["name"], params["email"]),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			service := Service{userStorer: test.userStorer}

			result := service.AddUser(test.params["name"], test.params["email"])

			require.Equal(t, test.expectedResponse, result)
		})
	}
}
