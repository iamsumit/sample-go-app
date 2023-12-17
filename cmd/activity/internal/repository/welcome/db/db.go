package welcomedb

// Repository holds the database logic for the router package.
type Repository struct{}

// New creates a new router repository.
func New() *Repository {
	return &Repository{}
}

// Welcome returns a welcome message.
func (s *Repository) Welcome(msg string) (string, error) {
	return msg, nil
}
