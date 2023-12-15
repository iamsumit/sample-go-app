package db

// Pagination structure to be used by the list endpoints.
//
// swagger:model list
type Pagination struct {
	// The current page
	//
	// in: query
	// type: integer
	// minimum: 1
	// required: false
	Page int `db:"page" json:"page" validate:"min=1"`

	// The per page limit
	//
	// in: query
	// type: integer
	// minimum: 1
	// maximum: 100
	// required: false
	PerPage int `db:"per_page" json:"per_page" validate:"min=1"`

	// The column to sort on
	//
	// in: query
	// required: false
	// type: string
	// enum: created,updated
	// description:
	//   Sort order:
	//   - `created` - When the record was created in the database
	//   - `updated` - When the record was last touched in the database
	Sort string `db:"sort" json:"sort" validate:"oneof=created updated"`

	// The direction of the sort
	//
	// in: query
	// required: false
	// type: string
	// enum: asc,desc
	// description:
	//   Sort order:
	//   - `asc` - Ascending, from A to Z
	//   - `desc` - Descending, from Z to A
	Direction string `db:"direction" json:"direction" validate:"oneof=asc desc"`
}

// Range returns the offset and limit for the pagination.
func (p Pagination) Range() (int, int) {
	// Calculate the offset.
	offset := (p.Page - 1) * p.PerPage
	if offset < 0 {
		offset = 0
	}

	return offset, p.PerPage
}

// SortDirection returns the sort and direction for the pagination.
//
// fields	- The map of fields from query to database field to sort on.
//
//	For example: map[string]string{"created": "created_at", "updated": "updated_at"}
func (p Pagination) SortDirection(fields map[string]string) (string, string) {
	for k, v := range fields {
		if p.Sort == k {
			p.Sort = v
		}
	}

	return p.Sort, p.Direction
}
