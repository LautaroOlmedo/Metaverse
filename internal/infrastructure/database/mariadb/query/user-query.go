package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"metaverse/internal/domain"
	"time"
)

const (
	getAllQuery        = `SELECT * FROM USERS`
	getOneQuery        = `SELECT * FROM USERS WHERE id= ?`
	getByUsernameQuery = ``
)

type UserQueryMariaDBRepository struct {
	db *sqlx.DB
}

func NewUserQueryMariaDBRepository(db *sqlx.DB) *UserQueryMariaDBRepository {
	return &UserQueryMariaDBRepository{
		db: db,
	}
}

func (repository *UserQueryMariaDBRepository) GetAllUsers(ctx context.Context, uChan chan []*domain.User, errChan chan error) {
	deadline := time.Now().Add(4 * time.Second)
	ctxD, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	doneChan := make(chan bool)

	go func() {
		<-ctxD.Done()
		doneChan <- true
		errChan <- errors.New("timeout")

	}()

	go func() {
		users := make([]*domain.User, 0)
		rows, err := repository.db.QueryContext(ctxD, getAllQuery)
		if err != nil {
			errChan <- err
			uChan <- nil
			doneChan <- true
			return
		}
		defer rows.Close()

		for rows.Next() {
			var u domain.User
			err := rows.Scan(&u.ID, &u.Name, &u.Email)
			if err != nil {
				errChan <- err
				uChan <- nil
				doneChan <- true
			}
			users = append(users, &u)
		}

		uChan <- users
		errChan <- nil
		doneChan <- true
	}()

	return

}

func (repository *UserQueryMariaDBRepository) GetOneUser(userID uuid.UUID, ctx context.Context, uChan chan *domain.User, errChan chan error) {
	deadline := time.Now().Add(4 * time.Second)
	ctxD, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	doneChan := make(chan bool)

	go func() {
		<-ctxD.Done()
		errChan <- fmt.Errorf("timeout")
		uChan <- nil
		doneChan <- true

	}()

	go func() {
		<-ctxD.Done()
		var u = &domain.User{}
		err := repository.db.QueryRowContext(ctx, getOneQuery, userID).Scan(u.ID, u.Name, u.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errChan <- fmt.Errorf("user not found")
			} else {
				errChan <- fmt.Errorf("internal server error")
			}
			uChan <- nil
			doneChan <- true
			return
		}
		uChan <- u
		errChan <- nil
		doneChan <- true

	}()
	return
}

func (repository *UserQueryMariaDBRepository) Login(username, password string) (*domain.User, error) {
	return nil, nil
}
