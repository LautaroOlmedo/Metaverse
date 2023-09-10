package domain

import (
	"github.com/google/uuid"
)

type Coin struct {
	ID          uuid.UUID
	PositionX   int8
	PositionY   int8
	PositionZ   int8
	IsCollected bool
}
