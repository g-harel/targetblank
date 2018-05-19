package page

// Group represents a group of items.
type Group struct {
	Items []*Item           `json:"items"`
	Meta  map[string]string `json:"meta"`
}

func newGroup() *Group {
	return &Group{
		Items: []*Item{},
		Meta:  map[string]string{},
	}
}
