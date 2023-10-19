package DTOs

import "github.com/google/uuid"

type UserDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" `
	Age      int8      `json:"age" `
	DNI      string    `json:"dni"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Active   bool      `json:"active"`
}
