package domain

import (
	"github.com/google/uuid"
	"metaverse/internal/domain/model"
)

type UserRepository interface {
	getAll() (map[uuid.UUID]model.User, error)
}
