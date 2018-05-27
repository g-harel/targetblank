package parser

import (
	"errors"
	"regexp"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Should return a syntax error when no rules match a line", func(t *testing.T) {
		p := New()
		err := p.Parse("test")
		if err == nil {
			t.Fatal("Expected lack of matching rules to produce error")
		}
		if strings.Index(err.Error(), "syntax") == -1 {
			t.Fatalf("Error does not contain \"syntax\": %v", err)
		}
	})

	t.Run("Should immediately return context errors and add line number", func(t *testing.T) {
		message := "test error"

		p := New()
		p.Add(Rule{
			Pattern: regexp.MustCompile(`a`),
			Handler: func(ctx *Context) {
				ctx.LineParsed()
			},
		}, Rule{
			Pattern: regexp.MustCompile(`b`),
			Handler: func(ctx *Context) {
				ctx.Error(message)
			},
		})

		err := p.Parse("b")
		if err == nil {
			t.Fatal("Expected parsing to produce error")
		}
		if strings.Index(err.Error(), message) == -1 {
			t.Fatalf("Expected context's error message to be returned: %v", err)
		}
		if err.Line != 1 {
			t.Fatalf("Expected correct line number")
		}

		err = p.Parse("a\na\nb")
		if err == nil {
			t.Fatal("Expected parsing to produce error")
		}
		if err.Line != 3 {
			t.Fatalf("Expected correct line number")
		}
	})

	t.Run("should run the rule handler when line matches pattern", func(t *testing.T) {
		s := "\n\n"
		count := 0

		p := New()
		p.Add(Rule{
			Pattern: regexp.MustCompile(`.*`),
			Handler: func(ctx *Context) {
				count++
				ctx.LineParsed()
			},
		})
		err := p.Parse(s)
		if err != nil {
			t.Fatalf("Unexpected parsing error: %v", err)
		}

		if count != len(strings.Split(s, "\n")) {
			t.Fatal("Handler was not called for the exact number of string lines")
		}
	})

	t.Run("should add named regexp groups to context", func(t *testing.T) {
		s := "a-bc\nde-f\nghi-jk"
		out := []string{}

		p := New()
		p.Add(Rule{
			Pattern: regexp.MustCompile(`^(?P<g1>\w+)-(?P<g2>\w+)$`),
			Handler: func(ctx *Context) {
				out = append(out, ctx.Param("g1")+"-"+ctx.Param("g2"))
				ctx.LineParsed()
			},
		})
		err := p.Parse(s)
		if err != nil {
			t.Fatalf("Unexpected parsing error: %v", err)
		}

		if strings.Join(out, "\n") != s {
			t.Fatalf("Reconstructed string is not equal to input (\"%v\" != \"%v\")", out, s)
		}
	})

	t.Run("should produce an error if a required rule does not match a line", func(t *testing.T) {
		s := "abc\ndef"

		p := New()
		p.Add(Rule{
			Pattern:  regexp.MustCompile(`abc`),
			Required: true,
			Handler: func(ctx *Context) {
				ctx.LineParsed()
			},
		})
		err := p.Parse(s)
		if err == nil {
			t.Fatal("Expected parsing error")
		}
	})

	t.Run("should produce an error if a required rule is not disabled after parsing", func(t *testing.T) {
		s := "abc\ndef"

		p := New()
		p.Add(Rule{
			Pattern:  regexp.MustCompile(`.*`),
			Required: true,
			Handler: func(ctx *Context) {
				ctx.LineParsed()
			},
		})
		err := p.Parse(s)
		if err == nil {
			t.Fatal("Expected parsing error")
		}
	})

	// TODO test more context funcs
}

func TestError(t *testing.T) {
	t.Run("Should add line number to error", func(t *testing.T) {
		err := Error{
			error: errors.New("error"),
			Line:  12,
		}
		if strings.Index(err.Error(), "12") == -1 {
			t.Fatalf("Expected line number to be in parser error message: %v", err)
		}
	})
}
