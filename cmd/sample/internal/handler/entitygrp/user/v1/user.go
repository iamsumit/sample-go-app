package user // import "github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/v1"

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

// Handler holds the dependencies for the user handler.
type Handler struct {
	Log logger.Logger
	DB  *sql.DB
}

// GetByID returns the user for the given id.
func (h *Handler) GetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := api.Param(r, "id")

	user := User{ID: id}
	api.Respond(ctx, w, user, http.StatusOK)

	return nil
}
