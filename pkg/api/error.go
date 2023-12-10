package api

import (
	"errors"
	"net/http"
)

//--------------------------------------------------------------------------
// Error Definitions
//--------------------------------------------------------------------------

var (
	// ErrDecode is used when there is an error decoding the request body.
	ErrDecode = NewError(
		errors.New("error decoding request body"),
		http.StatusBadRequest,
		nil,
	)
)
