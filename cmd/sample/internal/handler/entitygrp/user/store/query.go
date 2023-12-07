package store

import (
	"context"

	"github.com/iamsumit/sample-go-app/pkg/validator"
)

// Create creates a new user.
func (h *Handler) Create(ctx context.Context, user NewUser) (*User, error) {
	err := validator.Validate(user)
	if err != nil {
		return nil, err
	}

	return &User{}, nil
}
