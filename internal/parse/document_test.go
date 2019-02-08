package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func newDocument() *document {
	doc := &document{
		Meta:   map[string]string{},
		Groups: []*documentEntityGroup{},
	}
	doc.AddGroup()
	return doc
}

// Checks that the passed document and definition are equal when serialized to json.
func documentEquals(t *testing.T, target *document, s ...string) {
	raw := strings.Join(s, "\n")
	doc, err := Document(raw)
	if err != nil {
		t.Fatalf("Unexpected parsing error: %v\n%v", err, raw)
	}

	target.Raw = raw
	tb, err := json.MarshalIndent(target, "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling target document: %v", err)
	}
	tl := strings.Split(string(tb), "\n")

	rb := bytes.NewBuffer([]byte{})
	json.Indent(rb, []byte(doc), "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling result document: %v", err)
	}
	rl := strings.Split(rb.String(), "\n")

	// Helper for safe array access.
	getLine := func(n int, l []string) string {
		if n < 0 || len(l) <= n {
			return ""
		}
		return l[n]
	}

	// Loop through the maximum number of lines.
	for i := 0; i < len(tl) || i < len(rl); i++ {
		ts := getLine(i, tl)
		rs := getLine(i, rl)
		if rs != ts {
			// The diff is formatted to indicate the problematic line.
			ts = ">>" + (ts + "  ")[2:]
			rs = ">>" + (rs + "  ")[2:]

			// Two lines around the difference are also shown.
			t.Fatalf("Target and result documents do not match around line %v:\nEXPECTED: \n%v\nACTUAL: \n%v\n",
				i,
				strings.Join([]string{
					getLine(i-2, tl), getLine(i-1, tl), ts, getLine(i+1, tl), getLine(i+2, tl),
				}, "\n"),
				strings.Join([]string{
					getLine(i-2, rl), getLine(i-1, rl), rs, getLine(i+1, rl), getLine(i+2, rl),
				}, "\n"),
			)
		}
	}
}

// Checks assertions on expected parsing errors when parsing the passed document.
func produceErr(t *testing.T, line int, pattern string, s ...string) {
	raw := strings.Join(s, "\n")
	p, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("Error compiling expected error pattern /%v/: %v", pattern, err)
	}
	_, err = Document(raw)

	if err == nil || strings.Index(err.Error(), strconv.Itoa(line)) == -1 || !p.Match([]byte(err.Error())) {
		t.Fatalf("Parsing should have produced an error on line %v matching /%v/ but got: %v", line, pattern, err)
	}
}

func TestDocumentAddGroup(t *testing.T) {
	t.Run("should add a new group to the document", func(t *testing.T) {
		doc := newDocument()
		qty := len(doc.Groups)
		doc.AddGroup()

		if len(doc.Groups) <= qty {
			t.Errorf("Group was not added to document")
		}
	})
}

func TestDocumentAddGroupMeta(t *testing.T) {
	t.Run("should add values to last group's meta", func(t *testing.T) {
		doc := newDocument()
		key := "key"
		value := "value"
		doc.AddGroupMeta(key, value)

		if doc.Groups[0].Meta[key] != value {
			t.Errorf("Incorrect meta value (expected : \"%v\", received: \"%v\")", value, doc.Meta[key])
		}
	})
}

func TestDocumentEnter(t *testing.T) {
	t.Run("should add the entry to the last group when depth is zero", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			doc := newDocument()
			label := fmt.Sprintf("%d", i)

			for j := 0; j < i; j++ {
				doc.AddGroup()
			}

			err := doc.Enter(0, "link", label)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if doc.Groups[len(doc.Groups)-1].Entries[0].Label != label {
				t.Errorf("Item was not added to the correct group")
			}
		}
	})

	t.Run("should not allow negative depths", func(t *testing.T) {
		doc := newDocument()
		err := doc.Enter(-1, "", "")
		if err == nil {
			t.Fatalf("Expected invalid depth error")
		}
		if strings.Index(err.Error(), "invalid depth") == -1 {
			t.Fatalf("Expected invalid depth error, but got: %s", err)
		}
	})

	t.Run("should not allow depths to be skipped", func(t *testing.T) {
		doc := newDocument()
		err := doc.Enter(0, "", "")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		err = doc.Enter(2, "", "")
		if err == nil {
			t.Fatalf("Expected indentation level skipped error")
		}
		if strings.Index(err.Error(), "indentation level skipped") == -1 {
			t.Fatalf("Expected indentation level skipped error, but got: %s", err)
		}
	})

	t.Run("should trim whitespace from both link and label", func(t *testing.T) {
		link := " link "
		label := " label "
		doc := newDocument()
		doc.Enter(0, link, label)

		if doc.Groups[0].Entries[0].Link != strings.TrimSpace(link) {
			t.Errorf("Whitespace was not trimmed from :%s", link)
		}
		if doc.Groups[0].Entries[0].Label != strings.TrimSpace(label) {
			t.Errorf("Whitespace was not trimmed from :%s", label)
		}
	})

	t.Run("should use label as link if it resembles a url", func(t *testing.T) {
		expectURL := func(isURL bool, label string) {
			doc := newDocument()
			doc.Enter(0, "", label)
			if isURL {
				if doc.Groups[0].Entries[0].Link != label {
					t.Errorf("Label should have been used as link: \"%v\"", label)
				}
			} else {
				if doc.Groups[0].Entries[0].Link == label {
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

	t.Run("should correctly append entries", func(t *testing.T) {
		doc := newDocument()

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
			err := doc.Enter(depth, l, "")
			if err != nil {
				t.Fatalf("Unexpected error when adding Item{depth:%v,label:\"%v\"}: %v", depth, l, err)
			}

			current := doc.Groups[len(doc.Groups)-1].Entries[address[0]]
			for _, i := range address[1:] {
				if current.Children == nil || len(current.Children) <= i {
					t.Fatalf("Item was not added at specified address: %v", l)
				}
				current = current.Children[i]
			}

			if current.Label != l {
				t.Fatalf("Incorrect entry found at specified address (expected \"%v\", found: \"%v\")", l, current.Label)
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
		doc.AddGroup()       // ---
		check(0, 0)          // x
		check(0, 1)          // x
		check(1, 1, 0)       //  x
		check(2, 1, 0, 0)    //   x
		check(1, 1, 1)       //  x
		doc.AddGroup()       // ---
		check(0, 0)          // x
		check(1, 0, 0)       //  x
		check(1, 0, 1)       //  x
		check(1, 0, 2)       //  x
		check(1, 0, 3)       //  x
	})
}

func TestDocument(t *testing.T) {
	t.Run("should require version to be declared", func(t *testing.T) {
		produceErr(t, 2, "version",
			"",
		)
	})

	t.Run("should require a header", func(t *testing.T) {
		produceErr(t, 2, "header",
			"version 1",
		)
	})

	t.Run("should correctly parse the version from a minimal document", func(t *testing.T) {
		doc := newDocument()
		doc.Version = "1"
		documentEquals(t, doc,
			"version 1",
			"===",
		)
	})

	t.Run("should correctly ignore blank lines", func(t *testing.T) {
		doc := newDocument()
		doc.Version = "1"
		documentEquals(t, doc,
			"",
			"version 1",
			"",
			"===",
			"",
		)
	})

	t.Run("should correctly ignore trailing whitespace", func(t *testing.T) {
		doc := newDocument()
		doc.Version = "1"
		documentEquals(t, doc,
			"version 1 ",
			"===       ",
		)
	})

	t.Run("should correctly ignore comments", func(t *testing.T) {
		doc := newDocument()
		doc.Version = "1"
		documentEquals(t, doc,
			"          # comment",
			"version 1 # other comment",
			"===       # other other comment",
			"# commented line",
		)
	})

	t.Run("should not accept version > 1", func(t *testing.T) {
		produceErr(t, 1, "unsupported version",
			"version 2",
			"===",
		)
	})

	t.Run("v1", func(t *testing.T) {
		t.Run("should accept version 1", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			documentEquals(t, doc,
				"version 1",
				"===",
			)
		})

		t.Run("should correctly add documentEquals metadata", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.Meta["key1"] = "value1"
			doc.Meta["key2"] = "value2"
			doc.Meta["kEy3"] = "!@ $%^& *()[] +-_= <>?,.; ':|\\`~"
			documentEquals(t, doc,
				"version 1",
				"key1=value1",
				"key2 = value2",
				"kEy3= !@ $%^& *()[] +-_= <>?,.; ':|\\`~",
				"===",
			)
		})

		t.Run("should not recognize non-alphanumeric characters [^A-Za-z0-9] as documentEquals meta keys", func(t *testing.T) {
			produceErr(t, 2, ".*",
				"version 1",
				"!@key=test",
				"===",
			)
		})

		t.Run("should allow for new groups", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.AddGroup()
			doc.AddGroup()
			documentEquals(t, doc,
				"version 1",
				"===",
				"---",
				"---",
			)
		})

		t.Run("should correctly identify entry's link and label", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.Enter(0, "link", "label")
			doc.Enter(0, "", "label")
			doc.Enter(0, "link", "")
			doc.Enter(0, "", "la[bel")
			doc.Enter(0, "link.link", "")
			documentEquals(t, doc,
				"version 1",
				"===",
				"label [link]",
				"label",
				"[link]",
				"la[bel",
				"link.link",
			)
		})

		t.Run("should assign entries to correct groups", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.Enter(0, "", "group1")
			doc.AddGroup().Enter(0, "", "group2")
			doc.AddGroup().Enter(0, "", "group3")
			documentEquals(t, doc,
				"version 1",
				"===",
				"group1",
				"---",
				"group2",
				"---",
				"group3",
			)
		})

		t.Run("should correctly add group metadata", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.AddGroupMeta("key1", "value1")
			doc.AddGroup()
			doc.AddGroupMeta("key2", "value2")
			doc.AddGroupMeta("kEy3", "!@ $%^& *()[] +-_= <>?,.; ':|\\`~")
			documentEquals(t, doc,
				"version 1",
				"===",
				"key1=value1",
				"---",
				"key2 = value2",
				"kEy3= !@ $%^& *()[] +-_= <>?,.; ':|\\`~",
			)
		})

		t.Run("should not recognize non-alphanumeric characters [^A-Za-z0-9_-] as documentEquals group keys", func(t *testing.T) {
			produceErr(t, 2, ".*",
				"version 1",
				"!@key=test",
				"===",
			)
		})

		t.Run("should not accept group metadata after the first entry", func(t *testing.T) {
			produceErr(t, 7, "syntax",
				"version 1",
				"===",
				"key1=value1",
				"---",
				"key2=value2",
				"test [link]",
				"key3=value3",
			)
		})

		t.Run("should enforce indentation of four spaces", func(t *testing.T) {
			produceErr(t, 5, "indent.*(four|4).*space",
				"version 1",
				"===",
				"label",
				"label",
				"  label",
			)
		})

		t.Run("should not allow labels to skip indentation levels", func(t *testing.T) {
			produceErr(t, 6, "skipped",
				"version 1",
				"===",
				"label",
				"    label",
				"label",
				"        label",
				"    label",
			)
		})

		t.Run("should correctly parse complex entry hierarchies", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.Enter(0, "", "label")
			doc.Enter(1, "", "label")
			doc.Enter(2, "", "label")
			doc.Enter(3, "", "label")
			doc.Enter(4, "", "label")
			doc.Enter(2, "", "label")
			doc.Enter(2, "", "label")
			doc.Enter(3, "", "label")
			doc.Enter(4, "", "label")
			doc.Enter(4, "", "label")
			doc.Enter(0, "", "label")
			documentEquals(t, doc,
				"version 1",
				"===",
				"label",
				"    label",
				"        label",
				"            label",
				"                label",
				"        label",
				"        label",
				"            label",
				"                label",
				"                label",
				"label",
			)
		})

		t.Run("should correctly parse the original an example document", func(t *testing.T) {
			doc := newDocument()
			doc.Version = "1"
			doc.Meta["key"] = "value"
			doc.Meta["search"] = "google"
			doc.AddGroupMeta("key", "value")
			doc.Enter(0, "http://ee.co/1", "label_1")
			doc.Enter(0, "http://ee.co/2", "label 2")
			doc.Enter(1, "", "label3")
			doc.Enter(2, "http://ee.co/4", "label4")
			doc.Enter(1, "http://ee.co/5", "label-5")
			doc.AddGroup()
			doc.AddGroupMeta("name", "tasks")
			doc.Enter(0, "", "label6")
			doc.Enter(1, "", "label7")
			doc.Enter(1, "", "localhost:80/test")
			doc.Enter(1, "http://ee.co/10", "")
			doc.Enter(1, "", "label10")
			documentEquals(t, doc,
				"# single-line comments can be added anywhere",
				"version 1                       # version before any content",
				"                                # blank lines are ignored",
				"key=value                       # header contains key-value pairs",
				"search=google                   # ex. search bar provider is customizable",
				"===                             # header is mandatory",
				"key=value                       # groups can also have key-value pairs",
				"label_1 [http://ee.co/1]        # label can contain underscores",
				"label 2 [http://ee.co/2]        # label can contain spaces",
				"    label3                      # link is optional",
				"        label4 [http://ee.co/4] # list is infinitely nestable",
				"    label-5 [http://ee.co/5]    # label can contain dashes",
				"---                             # groups split layout into columns",
				"name=tasks                      # ex. group name can be added",
				"label6",
				"    label7                      # indentation level of 4 spaces",
				"    localhost:80/test           # labels that look like links should be clickable",
				"    [http://ee.co/10]           # label is optional",
				"    label10",
			)
		})
	})
}
