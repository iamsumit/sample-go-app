package activity

import (
	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/db"
)

// -------------------------------------------------------------------
// Swagger Parameters
// -------------------------------------------------------------------

// The parameters are applied to the create activity route.
//
// POST /v1/activity

// swagger:parameters createActivity
type _ struct {
	// The body to create a new activity.
	// in:body
	// required: true
	Body NewActivity `json:"body"`
}

// The parameters are applied to the get acvity route.
//
// GET /v1/activity/{id}

// swagger:parameters getActivity
type _ struct {
	// The id to get a new activity.
	// in:path
	// required: true
	// type: integer
	ID int `json:"id"`
}

// The parameters are applied to the list activity route.
//
// GET /v1/activities

// swagger:parameters listActivities
type _ struct {
	// The pagination to apply to the list activities.
	//
	// in: query
	// required: false
	db.Pagination `json:"page"`
}

// -------------------------------------------------------------------
// Swagger Responses
// -------------------------------------------------------------------

// Swagger response is used in various activity routes for a successful response.
//
// GET /v1/activity/{id}

// swagger:response activityResponse200
type _ struct {
	// in:body
	Body struct {
		// Success
		//
		// // example: false
		Success bool `json:"success"`
		// Timestamp
		//
		// example: 1639237536
		Timestamp int64 `json:"timestamp"`
		// Data
		// in: body
		Data []ActivityRes `json:"data"`
	}
}

// Swagger response is used in various activity routes for a successful response.
//
// POST /v1/activity

// swagger:response activityResponse201
type _ struct {
	// in:body
	Body struct {
		// Success
		//
		// // example: false
		Success bool `json:"success"`
		// Timestamp
		//
		// example: 1639237536
		Timestamp int64 `json:"timestamp"`
		// Data
		// in: body
		Data []ActivityRes `json:"data"`
	}
}

// Swagger response is used in various activity routes for a failed response
// because of 400.
//
// GET /v1/activity/{id}
// POST /v1/activity

// swagger:response activityResponse400
type _ struct {
	// in:body
	Body struct {
		// Success
		//
		// example: false
		Success bool `json:"success"`
		// Timestamp
		//
		// example: 1639237536
		Timestamp int64 `json:"timestamp"`
		// Data
		// in: body
		Errors api.ErrorResponse `json:"errors"`
	}
}

// Swagger response is used in various activity routes for a failed response
// because of 404.
//
// GET /v1/activity/{id}

// swagger:response activityResponse404
type _ struct {
	// in:body
	Body struct {
		// Success
		//
		// example: false
		Success bool `json:"success"`
		// Timestamp
		//
		// example: 1639237536
		Timestamp int64 `json:"timestamp"`
		// Data
		// in: body
		// example: {"error": "activity not found"}
		Errors map[string]interface{} `json:"errors"`
	}
}

// Swagger response is used in various activity routes for a failed response
// because of 500.
//
// GET /v1/activities
// GET /v1/activity/{id}
// POST /v1/activity

// swagger:response activityResponse500
type _ struct {
	// in:body
	Body struct {
		// Success
		//
		// example: false
		Success bool `json:"success"`
		// Timestamp
		//
		// example: 1639237536
		Timestamp int64 `json:"timestamp"`
		// Data
		// in: body
		// example: {"error": "some internal error occured"}
		Errors map[string]interface{} `json:"errors"`
	}
}
