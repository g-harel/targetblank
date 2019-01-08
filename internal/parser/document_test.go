package parser

import (
	"bytes"
	"encoding/json"
	"regexp"
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
	spec := strings.Join(s, "\n")
	result, parseErr := ParseDocument(spec)
	if parseErr != nil {
		t.Fatalf("Unexpected parsing error: %v\n%v", parseErr, spec)
	}

	target.Spec = spec
	tb, err := json.MarshalIndent(target, "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling target document: %v", err)
	}
	tl := strings.Split(string(tb), "\n")

	rb := bytes.NewBuffer([]byte{})
	json.Indent(rb, []byte(result), "| ", "  ")
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

// Checks assertions on expected parsing errors when parsing the passed spec.
func produceErr(t *testing.T, line int, pattern string, s ...string) {
	spec := strings.Join(s, "\n")
	p, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("Error compiling expected error pattern /%v/: %v", pattern, err)
	}
	_, parseErr := ParseDocument(spec)

	if parseErr == nil || parseErr.Line != line || !p.Match([]byte(parseErr.Error())) {
		t.Fatalf("Parsing should have produced an error on line %v matching /%v/ but got: %v", line, pattern, parseErr)
	}
}

func TestParser(t *testing.T) {
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

	t.Run("should correctly parse the version from a minimal spec", func(t *testing.T) {
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
			doc.Meta["kEy_-3"] = "!@ $%^& *()[] +-_= <>?,.; ':|\\`~"
			documentEquals(t, doc,
				"version 1",
				"key1=value1",
				"key2 = value2",
				"kEy_-3= !@ $%^& *()[] +-_= <>?,.; ':|\\`~",
				"===",
			)
		})

		t.Run("should not recognize non-alphanumeric characters [^A-Za-z0-9_-] as documentEquals meta keys", func(t *testing.T) {
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
			doc.AddGroupMeta("kEy_-3", "!@ $%^& *()[] +-_= <>?,.; ':|\\`~")
			documentEquals(t, doc,
				"version 1",
				"===",
				"key1=value1",
				"---",
				"key2 = value2",
				"kEy_-3= !@ $%^& *()[] +-_= <>?,.; ':|\\`~",
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
			doc := newDocument()
			doc.Version = "1"
			doc.AddGroupMeta("key1", "value1")
			doc.AddGroup()
			doc.AddGroupMeta("key2", "value2")
			doc.Enter(0, "link", "test")
			doc.Enter(0, "", "key3=value3")
			documentEquals(t, doc,
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
			produceErr(t, 6, "depth",
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

		t.Run("should correctly parse the original spec", func(t *testing.T) {
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
