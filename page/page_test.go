package page

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

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

func TestPageSetVersion(t *testing.T) {
	t.Run("should change the page version", func(t *testing.T) {
		p := New()
		version := "v1"
		p.SetVersion(version)

		if p.Version != version {
			t.Errorf("Incorrect version value (expected : \"%v\", received: \"%v\")", version, p.Version)
		}
	})
}

func TestPageAddMeta(t *testing.T) {
	t.Run("should add values to meta", func(t *testing.T) {
		p := New()
		key := "key"
		value := "value"
		p.AddMeta(key, value)

		if p.Meta[key] != value {
			t.Errorf("Incorrect meta value (expected : \"%v\", received: \"%v\")", value, p.Meta[key])
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

func TestPageAddItem(t *testing.T) {
	t.Run("should add the item to the last group when depth is zero", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			p := New()
			item := newItem("link", "label")

			for j := 0; j < i; j++ {
				p.AddGroup()
			}

			err := p.AddItem(0, item)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if p.Groups[len(p.Groups)-1].Items[0] != item {
				t.Errorf("Item was not added to the correct group")
			}
		}
	})

	t.Run("should not allow negative depths", func(t *testing.T) {
		p := New()
		err := p.AddItem(-1, newItem("", ""))
		if err == nil {
			t.Fatalf("Expected invalid depth error")
		}
		if strings.Index(err.Error(), "invalid depth") == -1 {
			t.Fatalf("Expected invalid depth error, but got: %s", err)
		}
	})

	t.Run("should not allow depths to be skipped", func(t *testing.T) {
		p := New()
		err := p.AddItem(0, newItem("", ""))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		err = p.AddItem(2, newItem("", ""))
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
			err := p.AddItem(depth, newItem(l, ""))
			if err != nil {
				t.Fatalf("Unexpected error when adding Item{depth:%v,label:\"%v\"}: %v", depth, l, err)
			}

			curr := p.Groups[len(p.Groups)-1].Items[address[0]]
			for _, i := range address[1:] {
				if curr.Items == nil || len(curr.Items) <= i {
					t.Fatalf("Item was not added at specified address: %v", l)
				}
				curr = curr.Items[i]
			}

			if curr.Label != l {
				t.Fatalf("Incorrect item found at specified address (expected \"%v\", found: \"%v\")", l, curr.Label)
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
