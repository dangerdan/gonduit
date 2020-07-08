package requests

import (
	"encoding/json"
	"errors"

	"github.com/dangerdan/gonduit/entities"
	"github.com/dangerdan/gonduit/util"
)

// DifferentialRevSearchRequest represents a request to differential.diff.search API method.
type DifferentialRevSearchRequest struct {
	// QueryKey is builtin or saved query to use. It is optional and sets initial constraints.
	QueryKey string `json:"queryKey,omitempty"`
	// Constraints contains additional filters for results. Applied on top of query if provided.
	Constraints *DifferentialRevSearchConstraints `json:"constraints,omitempty"`
	// Attachments specified what additional data should be returned with each result.
	Attachments *DifferentialRevSearchAttachments `json:"attachments,omitempty"`

	*entities.Cursor
	Request
}

// DifferentialRevSearchAttachments contains fields that specify what additional data should be returned with search results.
type DifferentialRevSearchAttachments struct {
	// Get information about subscribers.
	Subscribers bool `json:"subscribers,omitempty"`
	// Get the reviewers for each revision.
	Reviewers bool `json:"reviewers,omitempty"`
	// Get information about projects.
	Projects bool `json:"projects,omitempty"`
}

// DifferentialRevRequestSearchOrder describers how results should be ordered.
type DifferentialRevRequestSearchOrder struct {
	// Builtin is the name of predefined order to use.
	Builtin string
	// Order is list of columns to use for sorting, e.g. ["color", "-name", "id"],
	Order []string
}

// UnmarshalJSON parses JSON  into an instance of DifferentialRequestSearchOrder type.
func (o *DifferentialRevRequestSearchOrder) UnmarshalJSON(data []byte) error {
	if o == nil {
		return errors.New("differential search order is nil")
	}
	if jerr := json.Unmarshal(data, &o.Builtin); jerr == nil {
		return nil
	}

	return json.Unmarshal(data, &o.Order)
}

// MarshalJSON creates JSON our of DifferentialRequestSearchOrder instance.
func (o *DifferentialRevRequestSearchOrder) MarshalJSON() ([]byte, error) {
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

// DifferentialRevSearchConstraints describes search criteria for request.
type DifferentialRevSearchConstraints struct {
	// IDs - search for objects with specific IDs.
	IDs []int `json:"ids,omitempty"`
	// PHIDs - search for objects with specific PHIDs.
	PHIDs []string `json:"phids,omitempty"`
	// ResponsiblePHIDs - Find revisions that a given user is responsible for.
	ResponsiblePHIDs []string `json:"responsiblePHIDs,omitempty"`
	// Authors - Find revisions with specific authors.
	Authors []string `json:"authorPHIDs,omitempty"`
	// ReviewerPHIDs - Find revisions with specific reviewers.
	ReviewerPHIDs []string `json:"reviewerPHIDs,omitempty"`
	// RepositoryPHIDs - Find revisions from specific repositories.
	RepositoryPHIDs []string `json:"repositoryPHIDs,omitempty"`
	// Statuses - search for tasks with given statuses.
	Statuses []string `json:"statuses,omitempty"`
	// CreatedAfter - search for tasks created after given date.
	CreatedAfter *util.UnixTimestamp `json:"createdStart,omitempty"`
	// CreatedBefore - search for tasks created before given date.
	CreatedBefore *util.UnixTimestamp `json:"createdEnd,omitempty"`
	// ModifiedAfter - search for tasks modified after given date.
	ModifiedAfter *util.UnixTimestamp `json:"modifiedStart,omitempty"`
	// ModifiedBefore - search for tasks modified before given date.
	ModifiedBefore *util.UnixTimestamp `json:"modifiedEnd,omitempty"`
	// ClosedAfter - search for tasks closed after given date.
	Query string `json:"query,omitempty"`
	// Subscribers - search for objects with certain subscribers.
	Subscribers []string `json:"subscribers,omitempty"`
	// Projects - search for objects tagged with given projects.
	Projects []string `json:"projects,omitempty"`
}
