package welcome

// Service holds the business logic for the welcome package.
//
// This is an example of how to create a service interface.
type Service interface {
	Welcome(string) (string, error)
}

// WelcomeService holds the business logic for the welcome package.
//
// This is implementation of Service interface.
type WelcomeService struct {
	// repos is the repository for the welcome service.
	//
	// This could be a database repository or some third party API
	// that implements the interface.
	repos Repository
}

// NewService creates a new welcome service.
func NewService(r Repository) Service {
	return &WelcomeService{
		repos: r,
	}
}

// Welcome returns a welcome message.
func (s *WelcomeService) Welcome(message string) (string, error) {
	return s.repos.Welcome(message)
}
