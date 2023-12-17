package activity

import (
	"context"

	activityrepo "github.com/iamsumit/sample-go-app/activity/internal/repository/activity"
)

// Repository interface defines the methods that any repository should implement.
type Repository interface {
	Create(ctx context.Context, activity activityrepo.Activity) (*activityrepo.Activity, error)
}
