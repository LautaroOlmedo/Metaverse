package domain

import "github.com/google/uuid"

type MetaverseRoom struct {
	room  *Room
	users []*User
	coins []*Coin
}

func NewMetaverseRoom(xMin, xMax, yMin, yMax, zMin, zMax int8) (MetaverseRoom, error) {
	room := &Room{
		ID:   uuid.New(),
		Xmin: xMin,
		Xmax: xMax,
		Ymin: yMin,
		Ymax: yMax,
		Zmin: zMin,
		Zmax: zMax,
	}
	return MetaverseRoom{
		room:  room,
		users: make([]*User, 0),
		coins: make([]*Coin, 0),
	}, nil
}

func (mR *MetaverseRoom) GetID() uuid.UUID {
	return mR.room.ID
}
func (mR *MetaverseRoom) GetUsers() map[uuid.UUID]struct {
	Name           string
	Age            int8
	DNI            string
	Username       string
	Email          string
	CoinsCollected int16
	Position       struct {
		PositionX int16
		PositionY int16
		PositionZ int16
	}
} {

	usersMap := make(map[uuid.UUID]struct {
		Name           string
		Age            int8
		DNI            string
		Username       string
		Email          string
		CoinsCollected int16
		Position       struct {
			PositionX int16
			PositionY int16
			PositionZ int16
		}
	})

	for i := 0; i < len(mR.users); i++ {
		usersMap[mR.users[i].ID] = struct {
			Name           string
			Age            int8
			DNI            string
			Username       string
			Email          string
			CoinsCollected int16
			Position       struct {
				PositionX int16
				PositionY int16
				PositionZ int16
			}
		}{Name: mR.users[i].Name, Age: mR.users[i].Age, DNI: mR.users[i].DNI, Username: mR.users[i].Username, Email: mR.users[i].Email, CoinsCollected: mR.users[i].CoinsCollected, Position: struct {
			PositionX int16
			PositionY int16
			PositionZ int16
		}{PositionX: mR.users[i].Position.PositionX, PositionY: mR.users[i].Position.PositionY, PositionZ: mR.users[i].Position.PositionZ}}
	}
	return usersMap
}
