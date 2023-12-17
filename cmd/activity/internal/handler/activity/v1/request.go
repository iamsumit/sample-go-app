package activity

// -------------------------------------------------------------------
// Swagger Request models
// -------------------------------------------------------------------

// NewActivity represents the information required to create a New Activity.
//
// swagger:model createActivity
type NewActivity struct {
	// Name of the application to which this activity belongs
	//
	// in: body
	// type: string
	// required: true
	// example: sample
	AppName string `json:"app_name" mapstructure:"app_name" validate:"required,max=125"`

	// Name of the entity to which this activity belongs
	//
	// in: body
	// type: string
	// required: true
	// example: user
	Entity string `json:"entity" mapstructure:"entity" validate:"required,max=125"`

	// Operation performed on the entity
	//
	// in: body
	// type: string
	// required: true
	// example: created
	Operation string `json:"operation" mapstructure:"operation" validate:"required,max=125"`
}
