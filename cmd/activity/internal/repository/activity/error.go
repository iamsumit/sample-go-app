package activity

import (
	"errors"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/db"
)

var (
	// ErrAppNotFound is used when the activity is not found.
	ErrAppNotFound = db.NewError(
		errors.New("app not found"),
		http.StatusNotFound,
	)
)
