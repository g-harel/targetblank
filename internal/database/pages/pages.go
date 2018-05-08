package pages

import (
	"github.com/g-harel/targetblank/internal/database"
)

// Pages represents the table of page items.
type Pages struct {
	name   string
	client database.Client
}

// New creates a new Pages object.
func New(c database.Client) *Pages {
	return &Pages{
		name:   "targetblank-pages",
		client: c,
	}
}
