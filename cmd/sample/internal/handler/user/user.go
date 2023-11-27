package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"
)

type Handler struct {
	ctx context.Context
	log logger.Logger
	// @todo make it a wrapper on routing handler instead.
	tracer *tracer.Tracer
}

func New(ctx context.Context, log logger.Logger, tracer *tracer.Tracer) (*Handler, error) {
	return &Handler{
		ctx:    ctx,
		log:    log,
		tracer: tracer,
	}, nil
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	h.log.Info("User Handler entered")
	ctx, span := h.tracer.Start(r.Context(), "user.GetUser")
	span.SetAttributes(
		attribute.String("attribute-sample", "some-value"),
		attribute.String("url", r.URL.Path),
	)
	defer span.End()

	id := h.extractUIDFromRequest(ctx, r)

	h.log.Info("User Handler, ID: ", "id", id)
	fmt.Fprintf(w, "User: %v", id)
}

func (h *Handler) extractUIDFromRequest(ctx context.Context, r *http.Request) int {
	_, span := h.tracer.Start(ctx, "user.extractUIDFromRequest")
	span.SetAttributes(
		attribute.String("attribute-sample-child", "some-value"),
	)
	defer span.End()
	// Split the URL path by '/'
	parts := strings.Split(r.URL.Path, "/")

	// Get the last part of the URL path
	uid := parts[len(parts)-1]
	id, _ := strconv.Atoi(uid)

	if id < 4 {
		time.Sleep(time.Duration(id) * time.Second)
	}

	return id
}
