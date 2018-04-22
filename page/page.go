package page

import "fmt"

type Page struct {
	Version  string            `json:"version"`
	Meta     map[string]string `json:"meta"`
	Groups   []*Group          `json:"groups"`
	ancestry []*Item
}

func New() *Page {
	return &Page{
		Meta:     make(map[string]string),
		Groups:   []*Group{&Group{Items: []*Item{}}},
		ancestry: []*Item{},
	}
}

func (p *Page) SetVersion(v string) {
	p.Version = v
}

func (p *Page) AddMeta(key, value string) {
	p.Meta[key] = value
}

func (p *Page) AddGroup() {
	p.Groups = append(p.Groups, &Group{Items: []*Item{}})
	p.ancestry = []*Item{}
}

func (p *Page) AddItem(depth int, item *Item) error {
	if depth < 0 {
		return fmt.Errorf("invalid depth")
	}
	if depth == 0 {
		group := p.Groups[len(p.Groups)-1]
		group.Items = append(group.Items, item)
		p.ancestry = []*Item{item}
	} else if depth == len(p.ancestry) {
		parent := p.ancestry[depth-1]
		parent.Items = append(parent.Items, item)
		p.ancestry = append(p.ancestry, item)
	} else if depth < len(p.ancestry) {
		parent := p.ancestry[depth-1]
		parent.Items = append(parent.Items, item)
		p.ancestry = append(p.ancestry[depth-1:], item)
	} else {
		return fmt.Errorf("depth skipped")
	}
	return nil
}
