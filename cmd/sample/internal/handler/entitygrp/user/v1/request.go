// File: request.go
// -------------------------------------------------------------------
// Swagger Request models
// -------------------------------------------------------------------
package user

// NewUser represents the information required to create a New User.
//
// swagger:model createUser
type NewUser struct {
	// Name of the user
	//
	// in: body
	// type: string
	// required: true
	// example: Sumit Kumar
	Name string `json:"name"          mapstructure:"name"          validate:"required"`

	// the email address for this user
	//
	// in: body
	// type: string
	// example: user@provider.net
	Email *string `json:"email"         mapstructure:"email"         validate:"email"`

	// Bio of the user
	//
	// in: body
	// type: string
	// example: I am a developer by profession.
	Biography *string `json:"biography"     mapstructure:"biography"`

	// Date of birth of the user
	//
	// in: body
	// type: string
	// example: 1990-01-15
	DateOfBirth *string `json:"date_of_birth" mapstructure:"date_of_birth" validate:"birthDate"`
}
