package welcome

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds the endpoints for the welcome package.
type Endpoints struct {
	Welcome endpoint.Endpoint
}

// MakeEndpoints creates the endpoints for the welcome package.
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Welcome: WelcomeEndpoint(s),
	}
}

// WelcomeEndpoint returns the endpoint for the welcome handler.
func WelcomeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.Welcome("Welcome to Activity App!")
	}
}
