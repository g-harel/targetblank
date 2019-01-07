package mock

import (
	"fmt"

	"github.com/g-harel/targetblank/storage"
)

var pages = []*storage.Page{}

func PageCreate(p *storage.Page) (bool, error) {
	if p.Addr == "" {
		p.Addr = fmt.Sprintf("%06d", len(pages))
	}
	if p.Password == "" {
		p.Password = "tG6lUPO0OFxYFRgKaB2Cfts1UGdQX93w"
	}
	pages = append(pages, p)
	return false, nil
}

func PageRead(addr string) (*storage.Page, error) {
	for _, p := range pages {
		if p.Addr == addr {
			return p, nil
		}
	}
	return nil, nil
}

func PageUpdatePassword(addr, pass string) error {
	p, _ := PageRead(addr)
	p.Password = pass
	return nil
}

func PageUpdatePublished(addr string, published bool) error {
	p, _ := PageRead(addr)
	p.Published = published
	return nil
}

func PageUpdateData(addr, data string) error {
	p, _ := PageRead(addr)
	p.Data = data
	return nil
}

func PageDelete(addr string) error {
	for i, p := range pages {
		if p.Addr == addr {
			pages = append(pages[:i], pages[i+1:]...)
		}
	}
	return nil
}
