package activity

import (
	"time"

	activityrepo "github.com/iamsumit/sample-go-app/activity/internal/repository/activity"
)

// ActivityRes represents the information will be returned by the API.
//
// swagger:model activity
type ActivityRes struct {
	// ID of the activity
	//
	// type: int
	// required: true
	// example: 1
	ID int `json:"id"`

	// Entity of the activity
	//
	// type: string
	// required: true
	// example: user
	Entity string `json:"entity"`

	// Operation of the activity
	//
	// type: string
	// required: true
	// example: created
	Operation string `json:"operation"`

	// swagger:model app
	App App `json:"app"`

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

// App represents the application returned along with activity information.
//
// swagger:model app
type App struct {
	// ID of the app
	//
	// type: int
	// required: true
	// example: 1
	ID int `json:"id"`

	// Name of the app
	//
	// type: string
	// required: true
	// example: sample
	Name string `json:"name"`

	// Alias of the app
	//
	// type: string
	// required: true
	// example: sample
	Alias string `json:"alias"`

	// CreatedAt represents the time when the app was created.
	//
	// type: string
	// required: false
	// example: 2020-01-01T00:00:00Z
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt represents the time when the app was updated.
	//
	// type: string
	// required: false
	// example: 2020-01-01T00:00:00Z
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateFrom takes the store user and updates the response user.
func (a *ActivityRes) UpdateFrom(sa activityrepo.Activity) ActivityRes {
	a.ID = sa.ID
	a.Entity = sa.Entity
	a.Operation = sa.Operation
	a.CreatedAt = sa.CreatedAt
	a.UpdatedAt = sa.UpdatedAt

	a.App = App{
		ID:        sa.App.ID,
		Name:      sa.App.Name,
		Alias:     sa.App.Alias,
		CreatedAt: sa.App.CreatedAt,
		UpdatedAt: sa.App.UpdatedAt,
	}

	return *a
}
