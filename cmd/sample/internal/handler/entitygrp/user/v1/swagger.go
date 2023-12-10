// Package router Sample API.
//
// The purpose of this file is to define the swagger documentation for the API.
package user

import "github.com/iamsumit/sample-go-app/pkg/api"

// -------------------------------------------------------------------
// Swagger Parameters
// -------------------------------------------------------------------

// The parameters are applied to the create user route.
//
// POST /v1/user

// swagger:parameters createUser
type _ struct {
	// The body to create a new user.
	// in:body
	// required: true
	Body NewUser `json:"body"`
}

// The parameters are applied to the get user route.
//
// GET /v1/user/{id}

// swagger:parameters getUser
type _ struct {
	// The id to get a new user.
	// in:path
	// required: true
	// type: integer
	ID int `json:"id"`
}

// -------------------------------------------------------------------
// Swagger Responses
// -------------------------------------------------------------------

// Swagger response is used in various user routes for a successful response.
//
// GET /v1/user/{id}
// POST /v1/user

// swagger:response userResponse200
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
		Data []User `json:"data"`
	}
}

// Swagger response is used in various user routes for a failed response
// because of 404.
//
// GET /v1/user/{id}

// swagger:response userResponse404
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
		// example: {"error": "user not found"}
		Errors map[string]interface{} `json:"errors"`
	}
}

// Swagger response is used in various user routes for a failed response
// because of 400.
//
// GET /v1/user/{id}
// POST /v1/user

// swagger:response userResponse400
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
