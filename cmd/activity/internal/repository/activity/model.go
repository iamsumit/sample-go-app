// File: model.go declares the database request and response models used by the user store.
// -------------------------------------------------------------------
// Database request and response models
// -------------------------------------------------------------------
package activity

import "time"

// Activity represents an activity in the system.
//
// input for activity having input for app along with it.
type Activity struct {
	ID        int       `db:"id"`
	Entity    string    `db:"entity"`
	Operation string    `db:"operation"`
	App       App       `db:"-"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// App represents the service that activity belongs to.
//
// database input for app table, used in the app input.
type App struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Alias     string    `db:"alias"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
