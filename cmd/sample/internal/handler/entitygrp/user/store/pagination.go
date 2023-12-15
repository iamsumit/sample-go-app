package store

// Pagination holds the pagination parameters.
type Pagination struct {
	// Page is the page number.
	Page int `json:"page" validate:"required,min=1"`

	// PerPage is the number of items per page.
	PerPage int `json:"per_page" validate:"required,min=1,max=100"`

	// Sort is the name of the field to sort by.
	Sort string `json:"sort"`

	// Direction is the sort order.
	Direction string `json:"direction"`
}
