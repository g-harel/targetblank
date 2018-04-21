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

	emptyRule := parser.NewRule(
		"^\\s*$",
		func(ctx *parser.Context) {
			ctx.RemoveLine()
		},
	)
	ps.Add(emptyRule)

	whitespaceRule := parser.NewRule(
		"^(?P<content>.+?)\\s+$",
		func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	)
	ps.Add(whitespaceRule)

	commentRule := parser.NewRule(
		"^(?P<content>[^#]*)(#.*)$",
		func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	)
	ps.Add(commentRule)

	versionRule := parser.NewRule(
		"^(version:(?P<number>\\d+))?(?P<content>.*)$",
		func(ctx *parser.Context) {
			number := ctx.Param("number")
			if number == "" || ctx.Param("content") != "" {
				ctx.Error("could not find version")
				return
			}
			if number != "1" {
				ctx.Error("unsupported version")
				return
			}
			pg.Version = number
			ctx.RemoveLine()
			ctx.DisableRule()
		},
	)
	ps.Add(versionRule)

	metadataRule := parser.NewRule(
		"^(?P<key>[A-Za-z0-9_-]+)=\"(?P<value>.*)\"$",
		func(ctx *parser.Context) {
			pg.Meta[ctx.Param("key")] = ctx.Param("value")
			ctx.RemoveLine()
		},
	)
	ps.Add(metadataRule)

	headerRule := parser.NewRule(
		"^(===)?(?P<content>.*)$",
		func(ctx *parser.Context) {
			if ctx.Param("content") != "" {
				ctx.Error("could not parse header")
				return
			}
			ctx.RemoveLine()
			ctx.DisableRule()
		},
	)
	ps.Add(headerRule)

	currentGroup := 0
	prevIndentLevel := 0

	groupRule := parser.NewRule(
		"^---$",
		func(ctx *parser.Context) {
			currentGroup++
			pg.Groups = append(pg.Groups, &Group{Items: []*Item{}})
			prevIndentLevel = 0
			ctx.RemoveLine()
		},
	)
	ps.Add(groupRule)

	ancestry := []*Item{}
	labelRule := parser.NewRule(
		"^(?P<indent>(?:\\s{4})*)(?P<label>[^\\s].+?)?(?:\\[(?P<link>.*)\\])?$",
		func(ctx *parser.Context) {
			indent := ctx.Param("indent")
			label := strings.TrimSpace(ctx.Param("label"))
			link := strings.TrimSpace(ctx.Param("link"))

			indentLevel := 0
			if len(indent) > 0 {
				indentLevel = len(indent) / 4
			}
			if indentLevel-prevIndentLevel > 1 {
				ctx.Error("line over-indented")
				return
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
		},
	)
	ps.Add(labelRule)

	err := ps.Parse(spec)
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(pg, "", "    ")
	fmt.Println(string(b))
}
