package activity

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/validator"
)

// DecodeNewActivity decodes the new entity request object.
//
// It will be applicable to the request body of the following format:
//
//		{
//			"app_name": "sample",
//			"entity": "user",
//			"action": "create"
//	  }
func DecodeNewActivity(ctx context.Context, r *http.Request) (request interface{}, err error) {
	// -------------------------------------------------------------------
	// Decode the request object.
	// -------------------------------------------------------------------

	a := new(NewActivity)
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return nil, err
	}

	// -------------------------------------------------------------------
	// Validate the request object.
	// -------------------------------------------------------------------
	if err := validator.Validate(a); err != nil {
		return nil, err
	}

	return *a, nil
}
