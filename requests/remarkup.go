package requests

import "github.com/dangerdan/gonduit/constants"

// RemarkupProcessRequest represents a request to project.query.
type RemarkupProcessRequest struct {
	Context  constants.RemarkupProcessContextType `json:"context"`
	Contents []string                             `json:"contents"`
	Request
}
