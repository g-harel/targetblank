package page

import (
	"regexp"
	"strings"
)

// Matches strings that look like urls.
var urlPattern = regexp.MustCompile(`^\S+((\.|:)\S+)+(\/\S*)*$`)

// Item contains information about a single label/link and any potential children.
type Item struct {
	Label string  `json:"label"`
	Link  string  `json:"link"`
	Items []*Item `json:"items"`
}

func newItem(link, label string) *Item {
	link = strings.TrimSpace(link)
	label = strings.TrimSpace(label)

	if label == "" {
		label = link
	}
	if link == "" && urlPattern.Match([]byte(label)) {
		link = label
	}

	return &Item{
		Label: label,
		Link:  link,
		Items: []*Item{},
	}
}
