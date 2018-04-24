package page

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewItem(t *testing.T) {
	expectLink := func(t *testing.T, i *Item, expected string) {
		if i.Link != expected {
			t.Errorf("Incorrect link value (expected: \"%v\", received: \"%v\")", expected, i.Link)
		}
	}

	expectLabel := func(t *testing.T, i *Item, expected string) {
		if i.Label != expected {
			t.Errorf("Incorrect label value (expected: \"%v\", received: \"%v\")", expected, i.Label)
		}
	}

	t.Run("should use the provided link and label", func(t *testing.T) {
		link := "link"
		label := "label"
		i := NewItem(link, label)

		expectLink(t, i, link)
		expectLabel(t, i, label)
	})

	t.Run("should use the provided link as label if label is missing", func(t *testing.T) {
		link := "link"
		i := NewItem(link, "")

		expectLabel(t, i, link)
	})

	t.Run("should trim whitespace from both link and label", func(t *testing.T) {
		link := " link "
		label := " label "
		i := NewItem(link, label)

		expectLink(t, i, strings.TrimSpace(link))
		expectLabel(t, i, strings.TrimSpace(label))
	})

	t.Run("should create items that do not contain null values when serialized", func(t *testing.T) {
		i := NewItem("", "")
		b, _ := json.Marshal(i)
		s := string(b)

		if strings.Index(s, "null") >= 0 {
			t.Errorf("serialized object contains a null value: \"%v\"", s)
		}
	})

	t.Run("should use label as link if it resembles a url", func(t *testing.T) {
		expectURL := func(isURL bool, label string) {
			i := NewItem("", label)
			if isURL {
				if i.Link != label {
					t.Errorf("Label should have been used as link: \"%v\"", label)
				}
			} else {
				if i.Link == label {
					t.Errorf("Label should not have been used as link: \"%v\"", label)
				}
			}
		}

		expectURL(true, "https://example.com/test")
		expectURL(true, "www.example.com")
		expectURL(true, "example.com?q=test")
		expectURL(true, "localhost:8080")
		expectURL(false, "example: Example")
		expectURL(false, "ExampleExample")
		expectURL(false, "Example example example")
	})
}
