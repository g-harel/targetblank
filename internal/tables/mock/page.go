package mock

import (
	"errors"
	"fmt"

	"github.com/g-harel/targetblank/internal/tables"
)

// Unofficial count of created pages to generate unique addresses.
var count = 0

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
		count++
		item.Key = fmt.Sprintf("%08d", count)
	}
	if item.Password == "" {
		item.Password = "tG6lUPO0OFxYFRgKaB2Cfts1UGdQX93w"
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
