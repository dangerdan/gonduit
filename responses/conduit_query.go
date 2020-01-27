package responses

import "github.com/dangerdan/gonduit/entities"

// ConduitQueryResponse is the response of calling conduit.query.
type ConduitQueryResponse map[string]*entities.ConduitMethod
