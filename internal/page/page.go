package page

import "fmt"

// Page is a list of labels/links and metadata.
type Page struct {
	Version  string            `json:"version"`
	Meta     map[string]string `json:"meta"`
	Groups   []*Group          `json:"groups"`
	ancestry []*Item           // Maintains the list of parents of the last inserted item.
}

// New creates a new page with non-nil data structures.
func New() *Page {
	return &Page{
		Meta:     make(map[string]string),
		Groups:   []*Group{newGroup()},
		ancestry: []*Item{},
	}
}

// SetVersion changes the page's version.
func (p *Page) SetVersion(v string) *Page {
	p.Version = v
	return p
}

// AddMeta adds a key/value pair into the page's metadata map.
func (p *Page) AddMeta(key, value string) *Page {
	p.Meta[key] = value
	return p
}

// AddGroup adds a new empty group.
func (p *Page) AddGroup() *Page {
	p.Groups = append(p.Groups, newGroup())
	p.ancestry = []*Item{}
	return p
}

// AddGroupMeta adds a key/value pair into the last group's metadata map.
func (p *Page) AddGroupMeta(key, value string) *Page {
	p.Groups[len(p.Groups)-1].Meta[key] = value
	return p
}

// AddItem adds a new item to the page relative to the current ancestry.
func (p *Page) AddItem(depth int, item *Item) error {
	if depth == 0 {
		parent := p.Groups[len(p.Groups)-1]
		parent.Items = append(parent.Items, item)
		p.ancestry = []*Item{item}
		return nil
	}

	l := len(p.ancestry)
	if depth < 0 || depth > l {
		return fmt.Errorf("invalid depth")
	}

	if depth < l {
		p.ancestry = p.ancestry[:depth]
	}
	parent := p.ancestry[depth-1]
	parent.Items = append(parent.Items, item)
	p.ancestry = append(p.ancestry, item)
	return nil
}
