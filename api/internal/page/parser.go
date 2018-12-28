package page

import (
	"regexp"

	"github.com/g-harel/targetblank/api/internal/parser"
)

// NewFromSpec creates a new page from an input specification.
func NewFromSpec(s string) (*Page, *parser.Error) {
	p := New()
	p.Spec = s

	// v1HeaderMetadataRule matches with header metadata values.
	v1HeaderMetadataRule := parser.Rule{
		Name:     "header-metadata",
		Disabled: true,
		Pattern:  regexp.MustCompile(`^(?P<key>[A-Za-z0-9_-]+)\s*=\s*(?P<value>.*)$`),
		Handler: func(ctx *parser.Context) {
			p.Meta[ctx.Param("key")] = ctx.Param("value")
			ctx.LineParsed()
		},
	}

	// v1GroupMetadataRule matches with header metadata values.
	v1GroupMetadataRule := parser.Rule{
		Name:     "group-metadata",
		Disabled: true,
		Pattern:  v1HeaderMetadataRule.Pattern,
		Handler: func(ctx *parser.Context) {
			p.AddGroupMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.LineParsed()
		},
	}

	// v1GroupRule matches group delimiters.
	// These delimiters indicate a new group should be created.
	v1GroupRule := parser.Rule{
		Name:     "group",
		Disabled: true,
		Pattern:  regexp.MustCompile(`^---$`),
		Handler: func(ctx *parser.Context) {
			p.AddGroup()
			ctx.EnableOther(v1GroupMetadataRule.Name)
			ctx.LineParsed()
		},
	}

	// v1LabelRule matches labelled links.
	// Items are added to the page at the specified depth.
	v1LabelRule := parser.Rule{
		Name:     "label",
		Disabled: true,
		Pattern:  regexp.MustCompile(`^(?P<indent>\s*)(?P<label>[^\s\[].+?)?(?:\[(?P<link>.*)\])?$`),
		Handler: func(ctx *parser.Context) {
			indent := ctx.Param("indent")
			label := ctx.Param("label")
			link := ctx.Param("link")

			if len(indent)%4 != 0 {
				ctx.Error("expected indentation to be in 4 space increments")
				return
			}
			depth := len(indent) / 4

			err := p.AddItem(depth, newItem(link, label))
			if err != nil {
				ctx.Error(err.Error())
				return
			}

			ctx.DisableOther(v1GroupMetadataRule.Name)
			ctx.LineParsed()
		},
	}

	// v1HeaderRule is a required rule which matches with the header delimiter.
	// Once the header is found, the remaining rules are added to the parser.
	v1HeaderRule := parser.Rule{
		Name:     "header",
		Required: true,
		Pattern:  regexp.MustCompile(`^===$`),
		Handler: func(ctx *parser.Context) {
			ctx.DisableSelf()
			ctx.DisableOther(v1HeaderMetadataRule.Name)
			ctx.EnableOther(v1GroupMetadataRule.Name)
			ctx.EnableOther(v1GroupRule.Name)
			ctx.EnableOther(v1LabelRule.Name)
			ctx.LineParsed()
		},
	}

	// emptyRule removes lines that are entirely whitespace.
	emptyRule := parser.Rule{
		Name:    "empty",
		Pattern: regexp.MustCompile(`^\s*$`),
		Handler: func(ctx *parser.Context) {
			ctx.LineParsed()
		},
	}

	// whitespaceRule removes empty whitespace at the end of lines.
	whitespaceRule := parser.Rule{
		Name:    "whitespace",
		Pattern: regexp.MustCompile(`^(?P<content>.+?)\s+$`),
		Handler: func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	}

	// commentRule removes comments.
	commentRule := parser.Rule{
		Name:    "comment",
		Pattern: regexp.MustCompile(`^(?P<content>[^#]*)(#.*)$`),
		Handler: func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	}

	// versionRule is a required rule which matches with a version declaration.
	// When the version is found, the corresponding rules are added to the parser.
	versionRule := parser.Rule{
		Name:     "version",
		Required: true,
		Pattern:  regexp.MustCompile(`^version (?P<number>\d+)$`),
		Handler: func(ctx *parser.Context) {
			version := ctx.Param("number")
			if version == "1" {
				ctx.EnableOther(v1HeaderMetadataRule.Name)
				ctx.EnableOther(v1HeaderRule.Name)
			} else {
				ctx.Error("unsupported version")
				return
			}
			p.Version = version
			ctx.DisableSelf()
			ctx.LineParsed()
		},
	}

	ps := parser.New()
	ps.Add(
		emptyRule,
		whitespaceRule,
		commentRule,
		versionRule,
		v1HeaderMetadataRule,
		v1HeaderRule,
		v1GroupRule,
		v1GroupMetadataRule,
		v1LabelRule,
	)
	return p, ps.Parse(s)
}
