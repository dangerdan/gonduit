package responses

import "github.com/dangerdan/gonduit/entities"

// DifferentialRevisionSearchResponse is the response of calling differential.revision.search
type DifferentialRevisionSearchResponse struct {
	Data          map[string]entities.DiffusionCommit `json:"data"`
	IdentifierMap map[string]string                   `json:"identifierMap"`
	Cursor        entities.Cursor                     `json:"cursor"`
}
