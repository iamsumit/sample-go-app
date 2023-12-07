package user // import "github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/v1"

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

// Handler holds the dependencies for the user handler.
type Handler struct {
	log logger.Logger
	db  *sql.DB
}

func New(log logger.Logger, db *sql.DB) *Handler {
	return &Handler{
		log: log,
		db:  db,
	}
}

// GetByID returns the user for the given id.
func (h *Handler) GetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := api.Param(r, "id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.NewRequestError(errors.New("invalid id"), http.StatusBadRequest)
	}

	user := User{ID: uid}
	api.Respond(ctx, w, user, http.StatusOK)

	return nil
}

// CreateUser creates a new user.
func (h *Handler) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	user := new(User)
	err := api.Decode(r, user)
	if err != nil {
		h.log.Error(
			"decoding error",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)
		return api.NewRequestError(errors.New("unable to decode payload"), http.StatusBadRequest)
	}

	// @todo add database integration with proper validation in place.

	api.Respond(ctx, w, user, http.StatusOK)
	return nil
}
