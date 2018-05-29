package mock

import (
	"errors"

	"github.com/g-harel/targetblank/api/internal/rand"
	"github.com/g-harel/targetblank/api/internal/tables"
)

// Page is a mocked tables.Page object.
type Page struct {
	items []*tables.PageItem
}

// NewPage creates a new mocked tables.Page.
func NewPage() tables.IPage {
	return &Page{}
}

// Create adds a PageItem to the mocked Page table.
func (p *Page) Create(item *tables.PageItem) error {
	if item.Key == "" {
		item.Key = rand.String(6)
	}
	if item.Password == "" {
		item.Password = rand.String(32)
	}
	p.items = append(p.items, item)
	return nil
}

// Change modifies a PageItem in the mocked Page table.
func (p *Page) Change(addr string, i *tables.PageItem) error {
	for _, item := range p.items {
		if item.Key == addr {
			if i.Email != "" {
				item.Email = i.Email
			}
			if i.Password != "" {
				item.Password = i.Password
			}
			if i.PublishedHasBeenSetForUpdateExpression {
				item.Published = i.Published
			}
			if i.Page != "" {
				item.Page = i.Page
			}
			return nil
		}
	}
	return errors.New("item doesn't exist")
}

// Delete removes a PageItem from the mocked Page table.
func (p *Page) Delete(addr string) error {
	for i, item := range p.items {
		if item.Key == addr {
			p.items = append(p.items[:i], p.items[i+1:]...)
		}
	}
	return nil
}

// Fetch returns a PageItem from the mocked Page table.
func (p *Page) Fetch(addr string) (*tables.PageItem, error) {
	for _, item := range p.items {
		if item.Key == addr {
			return item, nil
		}
	}
	return nil, nil
}
