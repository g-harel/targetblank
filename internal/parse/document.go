package parse

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Document parses a document definition to JSON.
func Document(s string) (string, error) {
	doc := &document{
		Meta:   make(map[string]string),
		Groups: []*documentEntityGroup{},
	}
	doc.Spec = s
	doc.AddGroup()

	// v1HeaderMetadataRule matches with header metadata values.
	v1HeaderMetadataRule := rule{
		Name:     "header-metadata",
		Disabled: true,
		Pattern:  regexp.MustCompile(`^(?P<key>[A-Za-z0-9_-]+)\s*=\s*(?P<value>.*)$`),
		Handler: func(ctx *context) {
			doc.Meta[ctx.Param("key")] = ctx.Param("value")
			ctx.LineParsed()
		},
	}

	// v1GroupMetadataRule matches with header metadata values.
	v1GroupMetadataRule := rule{
		Name:     "group-metadata",
		Disabled: true,
		Pattern:  v1HeaderMetadataRule.Pattern,
		Handler: func(ctx *context) {
			doc.AddGroupMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.LineParsed()
		},
	}

	// v1GroupRule matches group delimiters.
	// These delimiters indicate a new group should be created.
	v1GroupRule := rule{
		Name:     "group",
		Disabled: true,
		Pattern:  regexp.MustCompile(`^---$`),
		Handler: func(ctx *context) {
			doc.AddGroup()
			ctx.EnableOther(v1GroupMetadataRule.Name)
			ctx.LineParsed()
		},
	}

	// v1EntryRule matches labelled links.
	// Items are added to the document at the specified depth.
	v1EntryRule := rule{
		Name:     "label",
		Disabled: true,
		Pattern:  regexp.MustCompile(`^(?P<indent>\s*)(?P<label>[^\s\[].+?)?(?:\[(?P<link>.*)\])?$`),
		Handler: func(ctx *context) {
			indent := ctx.Param("indent")
			label := ctx.Param("label")
			link := ctx.Param("link")

			if len(indent)%4 != 0 {
				ctx.Error("expected indentation to be in 4 space increments")
				return
			}
			depth := len(indent) / 4

			err := doc.Enter(depth, link, label)
			if err != nil {
				ctx.Error(err.Error())
				return
			}

			ctx.DisableOther(v1GroupMetadataRule.Name)
			ctx.LineParsed()
		},
	}

	// v1HeaderRule is a required rule which matches with the header delimiter.
	// Once the header is found, the remaining rules are added to the
	v1HeaderRule := rule{
		Name:     "header",
		Required: true,
		Pattern:  regexp.MustCompile(`^===$`),
		Handler: func(ctx *context) {
			ctx.DisableSelf()
			ctx.DisableOther(v1HeaderMetadataRule.Name)
			ctx.EnableOther(v1GroupMetadataRule.Name)
			ctx.EnableOther(v1GroupRule.Name)
			ctx.EnableOther(v1EntryRule.Name)
			ctx.LineParsed()
		},
	}

	// emptyRule removes lines that are entirely whitespace.
	emptyRule := rule{
		Name:    "empty",
		Pattern: regexp.MustCompile(`^\s*$`),
		Handler: func(ctx *context) {
			ctx.LineParsed()
		},
	}

	// whitespaceRule removes empty whitespace at the end of lines.
	whitespaceRule := rule{
		Name:    "whitespace",
		Pattern: regexp.MustCompile(`^(?P<content>.+?)\s+$`),
		Handler: func(ctx *context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	}

	// commentRule removes comments.
	commentRule := rule{
		Name:    "comment",
		Pattern: regexp.MustCompile(`^(?P<content>[^#]*)(#.*)$`),
		Handler: func(ctx *context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	}

	// versionRule is a required rule which matches with a version declaration.
	// When the version is found, the corresponding rules are added to the
	versionRule := rule{
		Name:     "version",
		Required: true,
		Pattern:  regexp.MustCompile(`^version (?P<number>\d+)$`),
		Handler: func(ctx *context) {
			version := ctx.Param("number")
			if version == "1" {
				ctx.EnableOther(v1HeaderMetadataRule.Name)
				ctx.EnableOther(v1HeaderRule.Name)
			} else {
				ctx.Error("unsupported version")
				return
			}
			doc.Version = version
			ctx.DisableSelf()
			ctx.LineParsed()
		},
	}

	err := (&parser{}).Add(
		emptyRule,
		whitespaceRule,
		commentRule,
		versionRule,
		v1HeaderMetadataRule,
		v1HeaderRule,
		v1GroupRule,
		v1GroupMetadataRule,
		v1EntryRule,
	).Parse(s)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(doc)
	if err != nil {
		return "", fmt.Errorf("marshal parse output: %v", err)
	}

	return string(b), nil
}

// Document represents document data and exposes helpers for the parsing rules.
type document struct {
	Version string                 `json:"version"`
	Spec    string                 `json:"spec"`
	Meta    map[string]string      `json:"meta"`
	Groups  []*documentEntityGroup `json:"groups"`
}

// DocumentEntityGroup represents a group of entries attached to a document.
type documentEntityGroup struct {
	Entries []*documentEntry  `json:"entries"`
	Meta    map[string]string `json:"meta"`
}

// DocumentEntry contains information about a single label/link and any potential children.
type documentEntry struct {
	Label    string           `json:"label"`
	Link     string           `json:"link"`
	Children []*documentEntry `json:"entries"`
}

// AddGroup adds a new empty group.
func (p *document) AddGroup() *document {
	p.Groups = append(p.Groups, &documentEntityGroup{
		Entries: []*documentEntry{},
		Meta:    map[string]string{},
	})
	return p
}

// AddGroupMeta adds a key/value pair into the last group's metadata map.
func (p *document) AddGroupMeta(key, value string) *document {
	p.Groups[len(p.Groups)-1].Meta[key] = value
	return p
}

// Enter adds a new entry relative to the most recent one.
func (p *document) Enter(depth int, link, label string) error {
	if depth < 0 {
		return fmt.Errorf("invalid depth")
	}

	link = strings.TrimSpace(link)
	label = strings.TrimSpace(label)
	if label == "" {
		label = link
	}
	match, err := regexp.MatchString(`^\S+((\.|:)\S+)+(\/\S*)*$`, label)
	if link == "" && err == nil && match {
		link = label
	}

	entry := &documentEntry{
		Label:    label,
		Link:     link,
		Children: []*documentEntry{},
	}

	parentGroup := p.Groups[len(p.Groups)-1]

	if depth == 0 {
		parentGroup.Entries = append(parentGroup.Entries, entry)
		return nil
	}

	if len(parentGroup.Entries) == 0 {
		return fmt.Errorf("invalid depth")
	}

	parent := parentGroup.Entries[len(parentGroup.Entries)-1]
	for i := 1; i < depth; i++ {
		if len(parent.Children) == 0 {
			return fmt.Errorf("invalid depth")
		}
		parent = parent.Children[len(parent.Children)-1]
	}

	parent.Children = append(parent.Children, entry)
	return nil
}
