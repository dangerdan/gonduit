package requests

// PHIDQueryRequest allows phids to be queried
type PHIDQueryRequest struct {
	PHIDs []string `json:"phids"`
	Request
}
