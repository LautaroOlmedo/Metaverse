package ports

import "context"

// UserQueryRepository is the interface that wraps the queries operations
//
//go:generate mockery --name=UserQueryRepository --output=domain --inpackage=true
type UserQueryRepository interface {
	Login(email, password string) error
}

// UserCommandRepository is the interface that wraps the commands operations
//
//go:generate mockery --name=UserCommandRepository --output=domain --inpackage=true
type UserCommandRepository interface {
	Register(ctx context.Context, id, name, dni, username, email, password string, age int8) error
}

// UserService is the business logic of the application
//
//go:generate mockery --name=UserService --output=domain --inpackage=true
type UserService interface {
	Login(email string, password string) error
	Register(email string, password string, passConfirm string) error
}

/*
// UserHandlers is the business logic of the infrastructure. REST-API
//
//go:generate mockery --name=Handler --output=services --inpackage=true
type UserHandlers interface{}*/
