package pages

import (
	"github.com/g-harel/targetblank/internal/database"
)

// Item represents a document stored in the page table.
type Item struct {
	Addr     string `json:"addr"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
	Temp     bool   `json:"temporary"`
	Page     string `json:"page"`
}

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
