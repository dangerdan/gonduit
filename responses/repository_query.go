package responses

import "github.com/dangerdan/gonduit/entities"

// RepositoryQueryResponse is the result of repository.query operations.
type RepositoryQueryResponse []*entities.Repository
