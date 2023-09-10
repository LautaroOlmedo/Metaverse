package model

import (
	"github.com/google/uuid"
	"metaverse/internal/domain"
)

type User struct {
	ID             uuid.UUID             `db:"id"`
	Name           string                `db:"name" `
	Age            int8                  `db:"age" `
	DNI            string                `db:"dni"`
	Username       string                `db:"username"`
	Email          string                `db:"email"`
	CoinsCollected int16                 `db:"coins_collected"`
	Position       domain.ActualPosition `db:"position"`
}
