package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/g-harel/targetblank/parser"
)

type Item struct {
	Label string  `json:"label"`
	Link  string  `json:"link"`
	Items []*Item `json:"items"`
}

type Group struct {
	Items []*Item `json:"items"`
}

type Page struct {
	Version string            `json:"version"`
	Meta    map[string]string `json:"meta"`
	Groups  []*Group          `json:"groups"`
}

func main() {
	pg := Page{
		Meta:   make(map[string]string),
		Groups: []*Group{&Group{Items: []*Item{}}},
	}
	ps := parser.New()

	// removes blank lines
	ps.AddNewRule(
		"^\\s*$",
		func(ctx *parser.Context, p map[string]string) error {
			ctx.RemoveLine()
			return nil
		},
	)

	// removes trailing whitespace
	ps.AddNewRule(
		"^(?P<content>.+?)\\s+$",
		func(ctx *parser.Context, p map[string]string) error {
			ctx.ReplaceLine(p["content"])
			return nil
		},
	)

	// removes comments
	ps.AddNewRule(
		"^(?P<content>[^#]*)(#.*)$",
		func(ctx *parser.Context, p map[string]string) error {
			ctx.ReplaceLine(p["content"])
			return nil
		},
	)

	// checks that a supported version is declared
	var versionRule *parser.Rule
	versionRule = parser.NewRule(
		"^(version:(?P<number>\\d+))?(?P<content>.*)$",
		func(ctx *parser.Context, p map[string]string) error {
			number := p["number"]
			if number == "" || p["content"] != "" {
				return fmt.Errorf("could not find version")
			}
			if number != "1" {
				return fmt.Errorf("unsupported version")
			}
			pg.Version = number
			ctx.RemoveLine()
			versionRule.Disable()
			return nil
		},
	)
	ps.Add(versionRule)

	// matches metadata key-value pairs
	ps.AddNewRule(
		"^(?P<key>[A-Za-z0-9_-]+)=\"(?P<value>.*)\"$",
		func(ctx *parser.Context, p map[string]string) error {
			pg.Meta[p["key"]] = p["value"]
			ctx.RemoveLine()
			return nil
		},
	)

	// matches the header
	var headerRule *parser.Rule
	headerRule = parser.NewRule(
		"^(===)?(?P<content>.*)$",
		func(ctx *parser.Context, p map[string]string) error {
			if p["content"] != "" {
				return fmt.Errorf("could not parse header")
			}
			ctx.RemoveLine()
			headerRule.Disable()
			return nil
		},
	)
	ps.Add(headerRule)

	currentGroup := 0
	prevIndentLevel := 0

	// matches group separators
	ps.AddNewRule(
		"^---$",
		func(ctx *parser.Context, p map[string]string) error {
			currentGroup++
			pg.Groups = append(pg.Groups, &Group{Items: []*Item{}})
			prevIndentLevel = 0
			ctx.RemoveLine()
			return nil
		},
	)

	// matches links and labels
	ancestry := []*Item{}
	ps.AddNewRule(
		"^(?P<indent>(?:\\s{4})*)(?P<label>[^\\s].+?)?(?:\\[(?P<link>.*)\\])?$",
		func(ctx *parser.Context, p map[string]string) error {
			indent := p["indent"]
			label := strings.TrimSpace(p["label"])
			link := strings.TrimSpace(p["link"])

			indentLevel := 0
			if len(indent) > 0 {
				indentLevel = len(indent) / 4
			}
			if indentLevel-prevIndentLevel > 1 {
				return fmt.Errorf("line over-indented")
			}

			if label == "" {
				label = link
			}
			if link == "" {
				link = label
			}

			item := &Item{
				Label: label,
				Link:  link,
				Items: []*Item{},
			}

			if indentLevel == 0 {
				pg.Groups[currentGroup].Items = append(
					pg.Groups[currentGroup].Items,
					item,
				)
				ancestry = []*Item{item}
			} else if indentLevel == prevIndentLevel {
				ancestry[len(ancestry)-2].Items = append(
					ancestry[len(ancestry)-2].Items,
					item,
				)
				ancestry[len(ancestry)-1] = item
			} else if indentLevel > prevIndentLevel {
				ancestry[len(ancestry)-1].Items = append(
					ancestry[len(ancestry)-1].Items,
					item,
				)
				ancestry = append(ancestry, item)
			} else {
				ancestry[indentLevel-1].Items = append(
					ancestry[indentLevel-1].Items,
					item,
				)
				ancestry = append(ancestry[indentLevel+1:], item)
			}

			prevIndentLevel = indentLevel

			ctx.RemoveLine()
			return nil
		},
	)

	err := ps.Parse(spec)
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(pg, "", "    ")
	fmt.Println(string(b))
}
