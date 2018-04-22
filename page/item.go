package page

type Item struct {
	Label string  `json:"label"`
	Link  string  `json:"link"`
	Items []*Item `json:"items"`
}

func NewItem(link, label string) *Item {
	if label == "" {
		label = link
	}
	if link == "" {
		link = label
	}

	return &Item{
		Label: label,
		Link:  link,
		Items: []*Item{},
	}
}
