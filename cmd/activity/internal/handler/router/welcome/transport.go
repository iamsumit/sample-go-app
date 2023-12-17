package welcome

import (
	"context"
	"net/http"
)

// Decode decodes the request object.
//
// Its an example of how to decode the request object.
func Decode(_ context.Context, r *http.Request) (request interface{}, err error) {
	// Decde your object here.
	return ExpectedInput{
		Message: "Welcome to Activity App!",
	}, nil
}
