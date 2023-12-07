package user

import (
	"time"

	"github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/store"
)

// User represents a user in the system.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Settings  Settings  `json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Settings represents the user settings.
type Settings struct {
	Email        *string    `json:"email"`
	Biography    *string    `json:"biography"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	IsActive     bool       `json:"is_active"`
	IsSubscribed bool       `json:"is_subscribed"`
}

// UpdateFrom takes the store user and updates the response user.
func (u *User) UpdateFrom(su store.User) User {
	u.ID = su.ID
	u.Name = su.Name
	u.Settings = Settings{
		Email:        su.Email,
		Biography:    su.Settings.Biography,
		DateOfBirth:  su.Settings.DateOfBirth,
		IsActive:     su.IsActive,
		IsSubscribed: su.Settings.IsSubscribed,
	}
	u.CreatedAt = su.CreatedAt
	u.UpdatedAt = su.UpdatedAt

	return *u
}
