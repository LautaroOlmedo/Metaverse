package services

import (
	"context"
	"errors"
	"metaverse/internal/core/ports"
	"os"
	"testing"
)

var userQuery *ports.MockUserQueryRepository
var userCommand *ports.MockUserCommandRepository

func TestMain(m *testing.M) {
	userQuery = &ports.MockUserQueryRepository{}
	userCommand = &ports.MockUserCommandRepository{}
	code := m.Run()
	os.Exit(code)
}
func TestGetAll(t *testing.T) {
	myContext := context.Background()

	type TestCase struct {
		test          string
		expectedError error
		email         string
		password      string
	}

	testCases := []TestCase{
		{
			test:          "PASS. valid user",
			email:         "lautaroolmedo77@gmail.com",
			password:      "validPassword",
			expectedError: nil,
		},
		{
			test:          "ERROR. invalid password",
			email:         "lautaroolmedo77@gmail.com",
			password:      "",
			expectedError: errors.New(""),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()

			err := userQuery.Login(myContext, tc.email, tc.password)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}
		})

	}
}
