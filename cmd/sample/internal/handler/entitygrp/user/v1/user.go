// Package user provides the handler for the user entity.
//
// It can be used to retrieve and store user information by given methods.
// It provides User model to be used to pass and retrieve user information.
package user // import "github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/v1"

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/validator"
	"github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/store"
)

// Handler holds the dependencies for the user handler.
type Handler struct {
	log   logger.Logger
	store *store.Handler
}

// New returns a new user handler for v1 version of user routes.
func New(log logger.Logger, db *sql.DB) *Handler {
	return &Handler{
		log:   log,
		store: store.New(db),
	}
}

// ByID returns the user for the given id.
func (h *Handler) ByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := api.Param(r, "id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	storeUser, err := h.store.ByID(ctx, uid)
	if err != nil && !errors.Is(err, store.ErrUserNotFound) {
		return err
	}

	if storeUser == nil {
		return api.NewRequestError(ErrUserNotFound, http.StatusNotFound)
	}

	user := new(User).UpdateFrom(*storeUser)
	err = api.Respond(ctx, w, user, http.StatusOK)

	return err
}

// CreateUser creates a new user.
func (h *Handler) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newUser := NewUser{}
	err := api.Decode(r, &newUser)
	if err != nil {
		h.log.Error(
			"decoding error",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return api.NewRequestError(ErrPayloadDecode, http.StatusBadRequest)
	}

	err = validator.Validate(newUser)
	if err != nil {
		return err
	}

	pTrue := true
	storeUser, err := h.store.Create(ctx, store.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		IsActive: &pTrue,
		Settings: store.Settings{
			IsSubscribed: &pTrue,
			Biography:    newUser.Biography,
			DateOfBirth:  newUser.DateOfBirth,
		},
	})
	if err != nil {
		return err
	}

	user := new(User).UpdateFrom(*storeUser)
	err = api.Respond(ctx, w, user, http.StatusOK)

	return err
}
