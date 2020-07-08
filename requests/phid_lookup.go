package requests

// PHIDLookupRequest provides PHID names
type PHIDLookupRequest struct {
	Names []string `json:"names"`
	Request
}
