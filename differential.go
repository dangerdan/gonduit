package gonduit

import (
	"github.com/dangerdan/gonduit/requests"
	"github.com/dangerdan/gonduit/responses"
)

// DifferentialQueryMethod is method name on Phabricator API.
const DifferentialQueryMethod = "differential.query"

// DifferentialQuery performs a call to differential.query.
func (c *Conn) DifferentialQuery(
	req requests.DifferentialQueryRequest,
) (*responses.DifferentialQueryResponse, error) {
	var res responses.DifferentialQueryResponse

	if err := c.Call(DifferentialQueryMethod, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialQueryDiffs performs a call to differential.querydiffs.
func (c *Conn) DifferentialQueryDiffs(
	req requests.DifferentialQueryDiffsRequest,
) (*responses.DifferentialQueryDiffsResponse, error) {
	var res responses.DifferentialQueryDiffsResponse

	if err := c.Call("differential.querydiffs", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialGetCommitPathsMethod is method name on Phabricator API.
const DifferentialGetCommitPathsMethod = "differential.getcommitpaths"

// DifferentialGetCommitPaths performs a call to differential.getcommitpaths.
func (c *Conn) DifferentialGetCommitPaths(
	req requests.DifferentialGetCommitPathsRequest,
) (*responses.DifferentialGetCommitPathsResponse, error) {
	var res responses.DifferentialGetCommitPathsResponse

	if err := c.Call(DifferentialGetCommitPathsMethod, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialGetCommitMessageMethod is method name on Phabricator API.
const DifferentialGetCommitMessageMethod = "differential.getcommitmessage"

// DifferentialGetCommitMessage performs a call to differential.getcommitmessage.
func (c *Conn) DifferentialGetCommitMessage(
	req requests.DifferentialGetCommitMessageRequest,
) (*responses.DifferentialGetCommitMessageResponse, error) {
	var res responses.DifferentialGetCommitMessageResponse

	if err := c.Call(
		DifferentialGetCommitMessageMethod, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
