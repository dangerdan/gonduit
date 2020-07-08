package responses

import "github.com/dangerdan/gonduit/entities"

// DifferentialDiffSearchResponse is the response of calling differential.diff.search.
type DifferentialDiffSearchResponse struct {
	Data          []entities.DifferentialDiffResult `json:"data"`
	IdentifierMap map[string]string                 `json:"identifierMap"`
	Cursor        entities.Cursor                   `json:"cursor"`
}
