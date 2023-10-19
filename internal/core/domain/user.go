package domain

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
)

var (
	InvalidEmail = errors.New("invalid email")
	InvalidName  = errors.New("the name field cannot be empty")
)

type User struct {
	id       uuid.UUID `db:"id"`
	name     string    `db:"name" `
	age      int8      `db:"age" `
	dni      string    `db:"dni"`
	username string    `db:"username"`
	email    string    `db:"email"`
	password string    `db:"password"`
	active   bool      `db:"active"`
}

// NewUser is Factory to create a new User
func NewUser(name, dni, username, email, password string, age int8) (User, error) {

	err := validateData(name, dni, username, email, password, age)
	if err != nil {
		return User{}, err
	}
	return User{
		id:       uuid.New(),
		name:     name,
		age:      age,
		dni:      dni,
		username: username,
		email:    email,
		password: password,
		active:   true,
	}, nil
}

func (u *User) GetID() string {
	return u.id.String()
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) UpdateEmail(newEmail string) {
	u.email = newEmail

}
func validateData(name, dni, username, email, password string, age int8) error {
	if !isEmailValid(email) {
		return InvalidEmail
	} else if name == "" {
		return InvalidName
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
