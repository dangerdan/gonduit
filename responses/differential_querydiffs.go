package responses

import "github.com/dangerdan/gonduit/entities"

// DifferentialQueryDiffsResponse is the response of calling differential.querydiffs.
type DifferentialQueryDiffsResponse []*entities.DifferentialDiff
