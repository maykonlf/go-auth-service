package entities

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
