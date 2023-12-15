// Package user provides the handler for the user entity.
//
// It can be used to retrieve and store user information by given methods.
// It provides User model to be used to pass and retrieve user information.
package user

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

// swagger:route GET /v1/users users listUsers
//
// All returns the list of users.
//
// This will help you get you list of the users from database.
//
// ---
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	Schemes: http, https
//
//	Responses:
//		- 200: userResponse200
//		- 500: userResponse500
func (h *Handler) All(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	pagi, err := api.PaginationParams(r)
	if err != nil {
		return err
	}

	storeUsers, err := h.store.All(ctx, pagi)
	if err != nil {
		return err
	}

	users := make([]User, len(storeUsers))
	for i, user := range storeUsers {
		users[i] = new(User).UpdateFrom(*user)
	}

	api.Respond(ctx, w, users, http.StatusOK)
	return nil
}

// swagger:route GET /v1/user/{id} users getUser
//
// ByID returns the user for the given id.
//
// This will help you get a user information by given id.
//
// ---
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	Schemes: http, https
//
//	Responses:
//		- 200: userResponse200
//		- 404: userResponse404
//	  - 400: userResponse400
//		- 500: userResponse500
func (h *Handler) ByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := api.Param(r, "id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.NewError(ErrInvalidID, http.StatusBadRequest, nil)
	}

	storeUser, err := h.store.ByID(ctx, uid)
	if err != nil && !errors.Is(err, store.ErrUserNotFound) {
		return err
	}

	if storeUser == nil {
		return api.NewError(ErrUserNotFound, http.StatusNotFound, nil)
	}

	user := new(User).UpdateFrom(*storeUser)
	err = api.Respond(ctx, w, []User{user}, http.StatusOK)

	return err
}

// swagger:route POST /v1/user users createUser
//
// Create a new user by given information.
//
// This will help you create a new user by given information.
// It will validate the information and create a new user.
// The uniqueness validation will be done if email is provided.
//
// ---
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	Schemes: http, https
//
//	Responses:
//		- 201: userResponse201
//		- 400: userResponse400
//	  - 409: userResponse409
//		- 500: userResponse500
func (h *Handler) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newUser := NewUser{}
	err := api.Decode(ctx, r, h.log, &newUser)
	if err != nil {
		return err
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
	err = api.Respond(ctx, w, []User{user}, http.StatusCreated)

	return err
}
