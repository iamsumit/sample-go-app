package activity

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds the endpoints for the activity package.
type Endpoints struct {
	Create endpoint.Endpoint
}

// MakeEndpoints creates the endpoints for the activity package.
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: Create(s),
	}
}

// swagger:route POST /v1/activity activity createActivity
//
// Create returns the endpoint for the activity handler.
//
// This will help you create a new activity by given information.
// It will validate the information and create a new activity.
//
// ---
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	Schemes: http, https
//
//	Responses:
//		- 201: activityResponse201
//	  - 400: activityResponse400
//		- 500: activityResponse500
func Create(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// -------------------------------------------------------------------
		// Process the activity creation request.
		// -------------------------------------------------------------------
		a, err := s.Create(ctx, request.(NewActivity))
		if err != nil {
			return nil, err
		}

		return a, nil
	}
}
