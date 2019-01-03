package page

import (
	"fmt"
	"regexp"
	"strings"
)

// Entry contains information about a single label/link and any potential children.
type Entry struct {
	Label    string   `json:"label"`
	Link     string   `json:"link"`
	Children []*Entry `json:"entries"`
}

// NewEntry creates a new entry and attempts to fill in the link or label if one is missing.
func newEntry(link, label string) *Entry {
	link = strings.TrimSpace(link)
	label = strings.TrimSpace(label)

	if label == "" {
		label = link
	}

	match, err := regexp.MatchString(`^\S+((\.|:)\S+)+(\/\S*)*$`, label)
	if link == "" && err == nil && match {
		link = label
	}

	return &Entry{
		Label:    label,
		Link:     link,
		Children: []*Entry{},
	}
}

// EntryGroup represents a group of entries.
type EntryGroup struct {
	Entries []*Entry          `json:"entries"`
	Meta    map[string]string `json:"meta"`
}

func newGroup() *EntryGroup {
	return &EntryGroup{
		Entries: []*Entry{},
		Meta:    map[string]string{},
	}
}

// Page is a collection of grouped entries and metadata.
type Page struct {
	Version string            `json:"version"`
	Spec    string            `json:"spec"`
	Meta    map[string]string `json:"meta"`
	Groups  []*EntryGroup     `json:"groups"`
}

// New creates a new page with non-nil data structures.
func New() *Page {
	p := &Page{
		Meta:   make(map[string]string),
		Groups: []*EntryGroup{},
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

// Enter adds a new item to the page relative.
func (p *Page) Enter(depth int, link, label string) error {
	if depth < 0 {
		return fmt.Errorf("invalid depth")
	}

	entry := newEntry(link, label)

	parentGroup := p.Groups[len(p.Groups)-1]

	if depth == 0 {
		parentGroup.Entries = append(parentGroup.Entries, entry)
		return nil
	}

	if len(parentGroup.Entries) == 0 {
		return fmt.Errorf("invalid depth")
	}

	parent := parentGroup.Entries[len(parentGroup.Entries)-1]
	for i := 1; i < depth; i++ {
		if len(parent.Children) == 0 {
			return fmt.Errorf("invalid depth")
		}
		parent = parent.Children[len(parent.Children)-1]
	}

	parent.Children = append(parent.Children, entry)
	return nil
}
