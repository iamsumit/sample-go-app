package store

import "time"

// User represents a user in the system.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	IsActive  bool      `json:"is_active"`
	Settings  Settings  `json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Settings represents the user settings.
type Settings struct {
	ID           int        `json:"id"`
	IsSubscribed bool       `json:"is_subscribed"`
	Biography    *string    `json:"biography"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
}

// NewUser represents the information required to create a New User.
type NewUser struct {
	Name        string     `json:"name" mapstructure:"name" validate:"required"`
	Email       *string    `json:"email" mapstructure:"email" validate:"email"`
	Biography   *string    `json:"biography" mapstructure:"biography"`
	DateOfBirth *time.Time `json:"date_of_birth" mapstructure:"date_of_birth"`
}
