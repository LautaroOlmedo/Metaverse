package domain

import (
	"context"
	"github.com/google/uuid"
)

// ReadUserRepository is the interface that wraps the queries operations
//
//go:generate mockery --name=Repository --output=domain --inpackage=true
type ReadUserRepository interface {
	GetAllUsers(ctx context.Context, uChan chan []*User, errChan chan error)
	GetOneUser(userID uuid.UUID, ctx context.Context, uChan chan *User, errChan chan error)
	Login(username, password string) (*User, error)
}

// WriteUserRepository is the interface that wraps the commands operations
//
//go:generate mockery --name=Repository --output=domain --inpackage=true
type WriteUserRepository interface {
	RegisterUser(name, email, username, password, dni string, age int8, position ActualPosition) error
	DeleteUser(userID uuid.UUID) error
	Delete(id int) error
}
