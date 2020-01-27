package gonduit

import (
	"github.com/dangerdan/gonduit/entities"
	"github.com/dangerdan/gonduit/requests"
	"github.com/dangerdan/gonduit/responses"
)

// ManiphestQuery performs a call to maniphest.query.
func (c *Conn) ManiphestQuery(
	req requests.ManiphestQueryRequest,
) (*responses.ManiphestQueryResponse, error) {
	var res responses.ManiphestQueryResponse

	if err := c.Call("maniphest.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ManiphestCreateTask performs a call to maniphest.createtask.
func (c *Conn) ManiphestCreateTask(
	req requests.ManiphestCreateTaskRequest,
) (*entities.ManiphestTask, error) {
	var res entities.ManiphestTask

	if err := c.Call("maniphest.createtask", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Conn) ManiphestGetTaskTransactions(
	req requests.ManiphestGetTaskTransactions,
) (*responses.ManiphestGetTaskTransactionsResponse, error) {
	var res responses.ManiphestGetTaskTransactionsResponse

	if err := c.Call("maniphest.gettasktransactions", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ManiphestSearch performs a call to maniphest.search.
func (c *Conn) ManiphestSearch(
	req requests.ManiphestSearchRequest,
) (*responses.ManiphestSearchResponse, error) {
	var res responses.ManiphestSearchResponse

	if err := c.Call("maniphest.search", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
