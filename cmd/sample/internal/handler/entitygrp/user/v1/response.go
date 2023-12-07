package user

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	IsActive  bool      `json:"is_active"`
	Settings  Settings  `json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Settings struct {
	ID           int        `json:"id"`
	IsSubscribed bool       `json:"is_subscribed"`
	Biography    *string    `json:"biography"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
}

// DecodeJSON decodes a JSON request body into the provided type.
func (u *User) DecodeJSON(data []byte) error {
	return json.Unmarshal(data, u)
}
