package command

import "github.com/jmoiron/sqlx"

type UserCommandMariaDBRepository struct {
	db *sqlx.DB
}

func NewUserCommandMariaDBRepository(db *sqlx.DB) *UserCommandMariaDBRepository {
	return &UserCommandMariaDBRepository{
		db: db,
	}
}
