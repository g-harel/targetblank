package page

import "fmt"

// Page is a list of labels/links and metadata.
type Page struct {
	Version string            `json:"version"`
	Spec    string            `json:"spec"`
	Meta    map[string]string `json:"meta"`
	Groups  []*Group          `json:"groups"`
}

// New creates a new page with non-nil data structures.
func New() *Page {
	p := &Page{
		Meta:   make(map[string]string),
		Groups: []*Group{},
	}
	return p.AddGroup()
}

// AddGroup adds a new empty group.
func (p *Page) AddGroup() *Page {
	p.Groups = append(p.Groups, newGroup())
	return p
}

// AddGroupMeta adds a key/value pair into the last group's metadata map.
func (p *Page) AddGroupMeta(key, value string) *Page {
	p.Groups[len(p.Groups)-1].Meta[key] = value
	return p
}

// AddItem adds a new item to the page relative to the current ancestry.
func (p *Page) AddItem(depth int, item *Item) error {
	if depth < 0 {
		return fmt.Errorf("invalid depth")
	}

	parentGroup := p.Groups[len(p.Groups)-1]

	if depth == 0 {
		parentGroup.Items = append(parentGroup.Items, item)
		return nil
	}

	if len(parentGroup.Items) == 0 {
		return fmt.Errorf("invalid depth")
	}

	parentItem := parentGroup.Items[len(parentGroup.Items)-1]
	for i := 1; i < depth; i++ {
		if len(parentItem.Items) == 0 {
			return fmt.Errorf("invalid depth")
		}
		parentItem = parentItem.Items[len(parentItem.Items)-1]
	}

	parentItem.Items = append(parentItem.Items, item)
	return nil
}
