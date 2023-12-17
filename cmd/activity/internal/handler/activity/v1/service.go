package activity

import (
	"context"

	activityrepo "github.com/iamsumit/sample-go-app/activity/internal/repository/activity"
)

// Service holds the business logic for the activity package.
//
// This is an example of how to create a service interface.
type Service interface {
	Create(ctx context.Context, activity NewActivity) (ActivityRes, error)
}

// ActivityService holds the business logic for the activity package.
//
// This is implementation of Service interface.
type ActivityService struct {
	// repos is the repository for the activity service.
	//
	// This could be a database repository or some third party API
	// that implements the interface.
	repos Repository
}

// NewService creates a new activity service.
func NewService(r Repository) Service {
	return &ActivityService{
		repos: r,
	}
}

// Create creates a new activity in the repository.
func (s *ActivityService) Create(ctx context.Context, activity NewActivity) (ActivityRes, error) {
	sa, err := s.repos.Create(ctx, activityrepo.Activity{
		Entity:    activity.Entity,
		Operation: activity.Operation,
		App: activityrepo.App{
			Name: activity.AppName,
		},
	})

	if sa == nil || err != nil {
		return ActivityRes{}, err
	}

	a := new(ActivityRes).UpdateFrom(*sa)

	return a, err
}
