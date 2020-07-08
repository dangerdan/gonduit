package requests

import (
	"encoding/json"
	"errors"

	"github.com/dangerdan/gonduit/entities"
)

// DifferentialDiffSearchRequest represents a request to differential.diff.search API method.
type DifferentialDiffSearchRequest struct {
	// QueryKey is builtin or saved query to use. It is optional and sets initial constraints.
	QueryKey string `json:"queryKey,omitempty"`
	// Constraints contains additional filters for results. Applied on top of query if provided.
	Constraints *DifferentialDiffSearchConstraints `json:"constraints,omitempty"`
	// Attachments specified what additional data should be returned with each result.
	Attachments *DifferentialDiffSearchAttachments `json:"attachments,omitempty"`
	// Cursor
	*entities.Cursor
	Request
}

// DifferentialDiffSearchAttachments contains fields that specify what additional data should be returned with search results.
type DifferentialDiffSearchAttachments struct {
	// Subscribers if true instructs server to return subscribers list for each task.
	Subscribers bool `json:"subscribers,omitempty"`
}

// DifferentialDiffRequestSearchOrder describers how results should be ordered.
type DifferentialDiffRequestSearchOrder struct {
	// Builtin is the name of predefined order to use.
	Builtin string
	// Order is list of columns to use for sorting, e.g. ["color", "-name", "id"],
	Order []string
}

// UnmarshalJSON parses JSON  into an instance of DifferentialRequestSearchOrder type.
func (o *DifferentialDiffRequestSearchOrder) UnmarshalJSON(data []byte) error {
	if o == nil {
		return errors.New("differential search order is nil")
	}
	if jerr := json.Unmarshal(data, &o.Builtin); jerr == nil {
		return nil
	}

	return json.Unmarshal(data, &o.Order)
}

// MarshalJSON creates JSON our of DifferentialRequestSearchOrder instance.
func (o *DifferentialDiffRequestSearchOrder) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nil, errors.New("differential search order is nil")
	}
	if o.Builtin != "" {
		return json.Marshal(o.Builtin)
	}
	if len(o.Order) > 0 {
		return json.Marshal(o.Order)
	}

	return nil, nil
}

// DifferentialDiffSearchConstraints describes search criteria for request.
type DifferentialDiffSearchConstraints struct {
	// IDs - search for objects with specific IDs.
	IDs []int `json:"ids,omitempty"`
	// PHIDs - Search for objects with specific PHIDs.
	PHIDs []string `json:"phids,omitempty"`
	// RevisionPHIDs - Find diffs attached to a particular revision.
	RevisionPHIDs []string `json:"revisionPHIDs,omitempty"`
	// Authors - search for tasks with given authors.
}
