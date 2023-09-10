package domain

import "github.com/google/uuid"

type Room struct {
	ID   uuid.UUID
	Xmin int8
	Xmax int8
	Ymin int8
	Ymax int8
	Zmin int8
	Zmax int8
}
