package responses

import "github.com/dangerdan/gonduit/entities"

// PHIDLookupResponse is the result of phid.lookup operations.
type PHIDLookupResponse map[string]*entities.PHIDResult
