package store

import "time"

// User represents a user in the system.
type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     *string   `db:"email"`
	IsActive  *bool     `db:"is_active"`
	Settings  Settings  `db:"-"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Settings represents the user settings.
type Settings struct {
	IsSubscribed *bool   `db:"is_subscribed"`
	Biography    *string `db:"biography"`
	DateOfBirth  *string `db:"date_of_birth"`
}
