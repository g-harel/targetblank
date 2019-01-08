package page

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"
)

// Checks that the passed page and the passed spec are equal when serialized to json.
func samePage(t *testing.T, target *Page, s ...string) {
	spec := strings.Join(s, "\n")
	result, perr := NewFromSpec(spec)
	if perr != nil {
		t.Fatalf("Unexpected parsing error: %v\n%v", perr, spec)
	}

	target.Spec = spec
	tb, err := json.MarshalIndent(target, "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling target Page: %v", err)
	}
	tl := strings.Split(string(tb), "\n")

	rb, err := json.MarshalIndent(result, "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling result Page: %v", err)
	}
	rl := strings.Split(string(rb), "\n")

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
			t.Fatalf("Target and result Pages do not match around line %v:\nEXPECTED: \n%v\nACTUAL: \n%v\n",
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
	_, perr := NewFromSpec(spec)

	if perr == nil || perr.Line != line || !p.Match([]byte(perr.Error())) {
		t.Fatalf("Parsing should have produced an error on line %v matching /%v/ but got: %v", line, pattern, perr)
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
		p := New()
		p.Version = "1"
		samePage(t, p,
			"version 1",
			"===",
		)
	})

	t.Run("should correctly ignore blank lines", func(t *testing.T) {
		p := New()
		p.Version = "1"
		samePage(t, p,
			"",
			"version 1",
			"",
			"===",
			"",
		)
	})

	t.Run("should correctly ignore trailing whitespace", func(t *testing.T) {
		p := New()
		p.Version = "1"
		samePage(t, p,
			"version 1 ",
			"===       ",
		)
	})

	t.Run("should correctly ignore comments", func(t *testing.T) {
		p := New()
		p.Version = "1"
		samePage(t, p,
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
			p := New()
			p.Version = "1"
			samePage(t, p,
				"version 1",
				"===",
			)
		})

		t.Run("should correctly add page metadata", func(t *testing.T) {
			p := New()
			p.Version = "1"
			p.Meta["key1"] = "value1"
			p.Meta["key2"] = "value2"
			p.Meta["kEy_-3"] = "!@ $%^& *()[] +-_= <>?,.; ':|\\`~"
			samePage(t, p,
				"version 1",
				"key1=value1",
				"key2 = value2",
				"kEy_-3= !@ $%^& *()[] +-_= <>?,.; ':|\\`~",
				"===",
			)
		})

		t.Run("should not recognize non-alphanumeric characters [^A-Za-z0-9_-] as page meta keys", func(t *testing.T) {
			produceErr(t, 2, ".*",
				"version 1",
				"!@key=test",
				"===",
			)
		})

		t.Run("should allow for new groups", func(t *testing.T) {
			p := New()
			p.Version = "1"
			p.AddGroup()
			p.AddGroup()
			samePage(t, p,
				"version 1",
				"===",
				"---",
				"---",
			)
		})

		t.Run("should correctly identify entry's link and label", func(t *testing.T) {
			p := New()
			p.Version = "1"
			p.Enter(0, "link", "label")
			p.Enter(0, "", "label")
			p.Enter(0, "link", "")
			p.Enter(0, "", "la[bel")
			p.Enter(0, "link.link", "")
			samePage(t, p,
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
			p := New()
			p.Version = "1"
			p.Enter(0, "", "group1")
			p.AddGroup().Enter(0, "", "group2")
			p.AddGroup().Enter(0, "", "group3")
			samePage(t, p,
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
			p := New()
			p.Version = "1"
			p.AddGroupMeta("key1", "value1")
			p.AddGroup()
			p.AddGroupMeta("key2", "value2")
			p.AddGroupMeta("kEy_-3", "!@ $%^& *()[] +-_= <>?,.; ':|\\`~")
			samePage(t, p,
				"version 1",
				"===",
				"key1=value1",
				"---",
				"key2 = value2",
				"kEy_-3= !@ $%^& *()[] +-_= <>?,.; ':|\\`~",
			)
		})

		t.Run("should not recognize non-alphanumeric characters [^A-Za-z0-9_-] as page group keys", func(t *testing.T) {
			produceErr(t, 2, ".*",
				"version 1",
				"!@key=test",
				"===",
			)
		})

		t.Run("should not accept group metadata after the first entry", func(t *testing.T) {
			p := New()
			p.Version = "1"
			p.AddGroupMeta("key1", "value1")
			p.AddGroup()
			p.AddGroupMeta("key2", "value2")
			p.Enter(0, "link", "test")
			p.Enter(0, "", "key3=value3")
			samePage(t, p,
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
			p := New()
			p.Version = "1"
			p.Enter(0, "", "label")
			p.Enter(1, "", "label")
			p.Enter(2, "", "label")
			p.Enter(3, "", "label")
			p.Enter(4, "", "label")
			p.Enter(2, "", "label")
			p.Enter(2, "", "label")
			p.Enter(3, "", "label")
			p.Enter(4, "", "label")
			p.Enter(4, "", "label")
			p.Enter(0, "", "label")
			samePage(t, p,
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
			p := New()
			p.Version = "1"
			p.Meta["key"] = "value"
			p.Meta["search"] = "google"
			p.AddGroupMeta("key", "value")
			p.Enter(0, "http://ee.co/1", "label_1")
			p.Enter(0, "http://ee.co/2", "label 2")
			p.Enter(1, "", "label3")
			p.Enter(2, "http://ee.co/4", "label4")
			p.Enter(1, "http://ee.co/5", "label-5")
			p.AddGroup()
			p.AddGroupMeta("name", "tasks")
			p.Enter(0, "", "label6")
			p.Enter(1, "", "label7")
			p.Enter(1, "", "localhost:80/test")
			p.Enter(1, "http://ee.co/10", "")
			p.Enter(1, "", "label10")
			samePage(t, p,
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
