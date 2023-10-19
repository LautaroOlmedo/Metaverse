package query

import (
	"github.com/jmoiron/sqlx"
)

const (
	queryLogin = `
         SELECT id, name, email, password FROM users WHERE email = ? AND password = ?;`
)

type UserQueryMariaDBRepository struct {
	db *sqlx.DB
}

func NewUserQueryMariaDBRepository(db *sqlx.DB) *UserQueryMariaDBRepository {
	return &UserQueryMariaDBRepository{
		db: db,
	}
}

func (queryRepo *UserQueryMariaDBRepository) Login(email, password string) error {
	//var u = &domain.User{}
	return nil
}
