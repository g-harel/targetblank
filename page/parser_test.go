package page

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"
)

func samePage(t *testing.T, target *Page, s ...string) {
	spec := strings.Join(s, "\n")
	result, perr := NewFromSpec(spec)
	if perr != nil {
		t.Fatalf("Unexpected parsing error: %v\n%v", perr, spec)
	}

	tb, err := json.MarshalIndent(target, "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling target Page: %v", err)
	}
	ts := string(tb)

	rb, err := json.MarshalIndent(result, "| ", "  ")
	if err != nil {
		t.Fatalf("Unexpected error when marshalling result Page: %v", err)
	}
	rs := string(rb)

	if ts != rs {
		t.Fatalf("Target and result Pages:\nEXPECTED: %v\nACTUAL: %v\n", ts, rs)
	}
}

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
		samePage(t, New().SetVersion("1"),
			"version 1",
			"===",
		)
	})

	t.Run("should correctly ignore blank lines", func(t *testing.T) {
		samePage(t, New().SetVersion("1"),
			"",
			"version 1",
			"",
			"===",
			"",
		)
	})

	t.Run("should correctly ignore trailing whitespace", func(t *testing.T) {
		samePage(t, New().SetVersion("1"),
			"version 1 ",
			"===       ",
		)
	})

	t.Run("should correctly ignore comments", func(t *testing.T) {
		samePage(t, New().SetVersion("1"),
			"          # comment",
			"version 1 # other comment",
			"===       # other other comment",
			"# commented line",
		)
	})

	t.Run("should correctly add metadata", func(t *testing.T) {
		samePage(t,
			New().
				SetVersion("1").
				AddMeta("key", "value"),
			"version 1",
			"key=value",
			"===",
		)
	})
}
