package responses

import "github.com/dangerdan/gonduit/entities"

// DifferentialQueryResponse is the response of calling differential.query.
type DifferentialQueryResponse []*entities.DifferentialRevision
