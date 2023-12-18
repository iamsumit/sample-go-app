package app

import "os"

var (
	// TraceName is the name of the trace.
	//
	// It can be used in different services so that everything can be tracked at one place.
	TraceName = "sample-go-app"
)

// Name returns the name of the application.
//
// It will be used as the service name in the traces.
// The name can be set using the APP_NAME environment variable.
// If not found, it will return the default value, "app".
func Name() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "app"
	}

	return appName
}
