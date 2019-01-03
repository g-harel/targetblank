package page

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestNewEntry(t *testing.T) {
	expectLink := func(t *testing.T, i *Entry, expected string) {
		if i.Link != expected {
			t.Errorf("Incorrect link value (expected: \"%v\", received: \"%v\")", expected, i.Link)
		}
	}

	expectLabel := func(t *testing.T, i *Entry, expected string) {
		if i.Label != expected {
			t.Errorf("Incorrect label value (expected: \"%v\", received: \"%v\")", expected, i.Label)
		}
	}

	t.Run("should use the provided link and label", func(t *testing.T) {
		link := "link"
		label := "label"
		i := newEntry(link, label)

		expectLink(t, i, link)
		expectLabel(t, i, label)
	})

	t.Run("should use the provided link as label if label is missing", func(t *testing.T) {
		link := "link"
		i := newEntry(link, "")

		expectLabel(t, i, link)
	})

	t.Run("should trim whitespace from both link and label", func(t *testing.T) {
		link := " link "
		label := " label "
		i := newEntry(link, label)

		expectLink(t, i, strings.TrimSpace(link))
		expectLabel(t, i, strings.TrimSpace(label))
	})

	t.Run("should create Items that do not contain null values when serialized", func(t *testing.T) {
		i := newEntry("", "")
		b, _ := json.Marshal(i)
		s := string(b)

		if strings.Index(s, "null") >= 0 {
			t.Errorf("Serialized Item contains a null value: \"%v\"", s)
		}
	})

	t.Run("should use label as link if it resembles a url", func(t *testing.T) {
		expectURL := func(isURL bool, label string) {
			i := newEntry("", label)
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

func TestNew(t *testing.T) {
	t.Run("should create Pages with an existing group", func(t *testing.T) {
		p := New()

		if len(p.Groups) == 0 {
			t.Errorf("New page does not have a group")
		}
	})

	t.Run("should create Pages that do not contain null values when serialized", func(t *testing.T) {
		p := New().AddGroup()
		b, _ := json.Marshal(p)
		s := string(b)

		if strings.Index(s, "null") >= 0 {
			t.Errorf("Serialized Page contains a null value: \"%v\"", s)
		}
	})
}

func TestPageAddGroup(t *testing.T) {
	t.Run("should add a new Group to the page", func(t *testing.T) {
		p := New()
		qty := len(p.Groups)
		p.AddGroup()

		if len(p.Groups) <= qty {
			t.Errorf("Group was not added to Page")
		}
	})
}

func TestPageAddGroupMeta(t *testing.T) {
	t.Run("should add values to last group's meta", func(t *testing.T) {
		p := New()
		key := "key"
		value := "value"
		p.AddGroupMeta(key, value)

		if p.Groups[0].Meta[key] != value {
			t.Errorf("Incorrect meta value (expected : \"%v\", received: \"%v\")", value, p.Meta[key])
		}
	})
}

func TestPageEnter(t *testing.T) {
	t.Run("should add the item to the last group when depth is zero", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			p := New()
			label := fmt.Sprintf("%d", i)

			for j := 0; j < i; j++ {
				p.AddGroup()
			}

			err := p.Enter(0, "link", label)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if p.Groups[len(p.Groups)-1].Entries[0].Label != label {
				t.Errorf("Item was not added to the correct group")
			}
		}
	})

	t.Run("should not allow negative depths", func(t *testing.T) {
		p := New()
		err := p.Enter(-1, "", "")
		if err == nil {
			t.Fatalf("Expected invalid depth error")
		}
		if strings.Index(err.Error(), "invalid depth") == -1 {
			t.Fatalf("Expected invalid depth error, but got: %s", err)
		}
	})

	t.Run("should not allow depths to be skipped", func(t *testing.T) {
		p := New()
		err := p.Enter(0, "", "")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		err = p.Enter(2, "", "")
		if err == nil {
			t.Fatalf("Expected invalid depth error")
		}
		if strings.Index(err.Error(), "invalid depth") == -1 {
			t.Fatalf("Expected invalid depth error, but got: %s", err)
		}
	})

	t.Run("should correctly append items", func(t *testing.T) {
		p := New()

		label := func(a []int) string {
			l := make([]string, len(a))
			for i, n := range a {
				l[i] = strconv.Itoa(n)
			}
			return strings.Join(l, ".")
		}

		check := func(depth int, address ...int) {
			if len(address) < 1 {
				t.Fatalf("Empty address passed as argument")
			}

			l := label(address)
			err := p.Enter(depth, l, "")
			if err != nil {
				t.Fatalf("Unexpected error when adding Item{depth:%v,label:\"%v\"}: %v", depth, l, err)
			}

			current := p.Groups[len(p.Groups)-1].Entries[address[0]]
			for _, i := range address[1:] {
				if current.Children == nil || len(current.Children) <= i {
					t.Fatalf("Item was not added at specified address: %v", l)
				}
				current = current.Children[i]
			}

			if current.Label != l {
				t.Fatalf("Incorrect item found at specified address (expected \"%v\", found: \"%v\")", l, current.Label)
			}
		}

		check(0, 0)          // x
		check(0, 1)          // x
		check(1, 1, 0)       //  x
		check(1, 1, 1)       //  x
		check(2, 1, 1, 0)    //   x
		check(1, 1, 2)       //  x
		check(2, 1, 2, 0)    //   x
		check(3, 1, 2, 0, 0) //    x
		check(3, 1, 2, 0, 1) //    x
		check(0, 2)          // x
		p.AddGroup()         // ---
		check(0, 0)          // x
		check(0, 1)          // x
		check(1, 1, 0)       //  x
		check(2, 1, 0, 0)    //   x
		check(1, 1, 1)       //  x
		p.AddGroup()         // ---
		check(0, 0)          // x
		check(1, 0, 0)       //  x
		check(1, 0, 1)       //  x
		check(1, 0, 2)       //  x
		check(1, 0, 3)       //  x
	})
}
