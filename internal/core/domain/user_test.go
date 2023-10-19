package domain

import (
	"errors"
	"github.com/google/uuid"
	"testing"
)

func Test_NewUser(t *testing.T) {
	type testCase struct {
		test          string
		id            uuid.UUID
		name          string
		age           int8
		dni           string
		username      string
		email         string
		password      string
		active        bool
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "PASS. Valid parameters",
			id:            uuid.New(),
			name:          "Lautaro",
			age:           23,
			dni:           "42509723",
			username:      "Chasx",
			email:         "lautaroolmedo77@gmail.com",
			password:      "secret",
			active:        true,
			expectedError: nil,
		},
		{
			test:          "ERROR. Invalid email",
			id:            uuid.New(),
			name:          "Joel",
			age:           30,
			dni:           "77777777",
			username:      "GermenDinger",
			email:         "germerDinger$gmail.com",
			password:      "noCarreo",
			active:        true,
			expectedError: InvalidEmail,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			_, err := NewUser(tc.name, tc.dni, tc.username, tc.email, tc.password, tc.age)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected: %v, got: %v", tc.expectedError, err)
			}

		})
	}
}
