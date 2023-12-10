// File: response.go
// -------------------------------------------------------------------
// Swagger Response models
// -------------------------------------------------------------------
package user

import (
	"time"

	"github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/store"
)

// User represents the information will be returned by the API.
//
// swagger:model user
type User struct {
	// ID of the user
	//
	// type: int
	// required: true
	// example: 1
	ID int `json:"id"`

	// Name of the user
	//
	// type: string
	// required: false
	// example: Sumit Kumar
	Name string `json:"name"`

	// swagger:model userSettings
	Settings Settings `json:"settings"`

	// CreatedAt represents the time when the user was created.
	//
	// type: string
	// required: false
	// example: 2020-01-01T00:00:00Z
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt represents the time when the user was updated.
	//
	// type: string
	// required: false
	// example: 2020-01-01T00:00:00Z
	UpdatedAt time.Time `json:"updated_at"`
}

// Settings represents the user settings returned along with user information.
//
// swagger:model userSettings
type Settings struct {
	// Email of the user
	//
	// type: string
	// required: false
	// example: user@provider.net
	Email *string `json:"email"`

	// Bio of the user
	//
	// type: string
	// required: false
	// example: I am a developer by profession.
	Biography *string `json:"biography"`

	// Date of birth of the user
	//
	// type: string
	// required: false
	// example: 1990-01-15
	DateOfBirth *string `json:"date_of_birth"`

	// IsActive represents the status of the user.
	//
	// type: bool
	// required: false
	// example: true
	IsActive *bool `json:"is_active"`

	// IsSubscribed represents the subscription status of the user.
	//
	// type: bool
	// required: false
	// example: true
	IsSubscribed *bool `json:"is_subscribed"`
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
