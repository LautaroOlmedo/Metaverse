package command

import (
	"context"
	"github.com/jmoiron/sqlx"
)

const (
	queryListUsers = `
         SELECT * FROM users;`
	queryInsertUser = `
         INSERT INTO users (id, name, email, username, password, dni, age)
         VALUES (?, ?, ?, ?, ?, ?, ?);`

	queryGetUserByEmail = `
         SELECT id, name, email, password FROM users WHERE email = ?;`

	queryGetUSer = `
         SELECT id, name, email FROM users WHERE id = ?;`
)

type UserCommandMariaDBRepository struct {
	db *sqlx.DB
}

func NewUserCommandMariaDBRepository(db *sqlx.DB) *UserCommandMariaDBRepository {
	return &UserCommandMariaDBRepository{
		db: db,
	}
}

func (commandRepo *UserCommandMariaDBRepository) Register(ctx context.Context, id, name, dni, username, email, password string, age int8) error {
	tx, err := commandRepo.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, queryInsertUser, id, name, email, username, password, dni, age)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
