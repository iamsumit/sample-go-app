package welcome

// Repository is an example of the interface that represents the repository.
//
// This is an example of how to create a repository interface.
type Repository interface {
	Welcome(string) (string, error)
}
