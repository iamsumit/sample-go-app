package user

// NewUser represents the information required to create a New User.
type NewUser struct {
	Name        string  `json:"name"          mapstructure:"name"          validate:"required"`
	Email       *string `json:"email"         mapstructure:"email"         validate:"email"`
	Biography   *string `json:"biography"     mapstructure:"biography"`
	DateOfBirth *string `json:"date_of_birth" mapstructure:"date_of_birth" validate:"birthDate"`
}
