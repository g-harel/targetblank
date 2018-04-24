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
		Groups:   []*Group{&Group{Items: []*Item{}}},
		ancestry: []*Item{},
	}
}

// SetVersion changes the page's version.
func (p *Page) SetVersion(v string) {
	p.Version = v
}

// AddMeta adds a key/value pair into the page's metadata map.
func (p *Page) AddMeta(key, value string) {
	p.Meta[key] = value
}

// AddGroup creates a new empty group.
func (p *Page) AddGroup() {
	p.Groups = append(p.Groups, &Group{Items: []*Item{}})
	p.ancestry = []*Item{}
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
