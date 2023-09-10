package domain

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID      `db:"id"`
	Name           string         `db:"name" `
	Age            int8           `db:"age" `
	DNI            string         `db:"dni"`
	Username       string         `db:"username"`
	Email          string         `db:"email"`
	Password       string         `db:"password"`
	CoinsCollected int16          `db:"coins_collected"`
	Position       ActualPosition `db:"position"`
}

type ActualPosition struct {
	PositionX int16
	PositionY int16
	PositionZ int16
}
