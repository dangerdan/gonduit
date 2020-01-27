package responses

import "github.com/dangerdan/gonduit/entities"

// PasteQueryResponse represents the result of calling paste.query.
type PasteQueryResponse map[string]*entities.PasteItem
