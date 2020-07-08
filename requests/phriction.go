package requests

// PhrictionInfoRequest allows requests via slug
type PhrictionInfoRequest struct {
	Slug string `json:"slug"`
	Request
}
