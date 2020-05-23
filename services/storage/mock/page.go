package mock

import (
	"fmt"

	"github.com/g-harel/targetblank/services/storage"
)

// Pages stores the test pages being added and removed.
var pages = []*storage.Page{}

// PageCreate adds a page to the internal store.
func PageCreate(p *storage.Page) (bool, error) {
	if p.Addr == "" {
		p.Addr = fmt.Sprintf("%06d", len(pages))
	}
	if p.Password == "" {
		p.Password = fmt.Sprintf("password-%06d", len(pages))
	}

	page, _ := PageRead(p.Addr)
	if page != nil {
		return true, nil
	}

	pages = append(pages, p)
	return false, nil
}

// PageRead reads a page from the internal store.
func PageRead(addr string) (*storage.Page, error) {
	for _, p := range pages {
		if p.Addr == addr {
			return p, nil
		}
	}
	return nil, nil
}

// PageUpdatePassword updates the password of a page in the internal store.
func PageUpdatePassword(addr, pass string) error {
	p, _ := PageRead(addr)
	p.Password = pass
	return nil
}

// PageUpdateDocument updates the document of a page in the internal store.
func PageUpdateDocument(addr, document string) error {
	p, _ := PageRead(addr)
	p.Document = document
	return nil
}

// PageDelete removes a page from the internal store.
func PageDelete(addr string) error {
	for i, p := range pages {
		if p.Addr == addr {
			pages = append(pages[:i], pages[i+1:]...)
		}
	}
	return nil
}
