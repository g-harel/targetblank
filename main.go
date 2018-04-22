package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

var spec = `
# single-line comments can be added anywhere
version 1                       # version before any content
                                # blank lines are ignored
key="value"                     # header contains key-value pairs
search="google"                 # search bar provider is customizable
===                             # header is mandatory
label_1 [http://ee.co/1]        # label can contain underscores
label 2 [http://ee.co/2]        # label can contain spaces
    label3                      # link is optional
        label4 [http://ee.co/4] # list is infinitely nestable
    label-5 [http://ee.co/5]    # label can contain dashes
---                             # groups split layout into columns
label6
    label7                      # indentation level of 4 spaces
    http://ee.co/9              # labels that look like links should be clickable
    [http://ee.co/10]           # label is optional
    label10
`

func main() {
	tmp := page.New()

	emptyRule := parser.NewRule().
		Pattern("\\s*").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
		})

	whitespaceRule := parser.NewRule().
		Pattern("(?P<content>.+?)\\s+").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})

	commentRule := parser.NewRule().
		Pattern("(?P<content>[^#]*)(#.*)").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})

	versionRule := parser.NewRule().
		Strict("unkown syntax is not version (version X)").
		Pattern("version (?P<number>\\d+)").
		Handler(func(ctx *parser.Context) {
			number := ctx.Param("number")
			if number != "1" {
				ctx.Error("unsupported version")
				return
			}
			tmp.SetVersion(number)
			ctx.IgnoreLine()
			ctx.DisableRule()
		})

	metadataRule := parser.NewRule().
		Pattern("(?P<key>[A-Za-z0-9_-]+)=\"(?P<value>.*)\"").
		Handler(func(ctx *parser.Context) {
			tmp.AddMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.IgnoreLine()
		})

	headerRule := parser.NewRule().
		Strict("unkown syntax is not header (===)").
		Pattern("===").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
			ctx.DisableRule()
			metadataRule.Disable()
		})

	groupRule := parser.NewRule().
		Pattern("---").
		Handler(func(ctx *parser.Context) {
			tmp.AddGroup()
			ctx.IgnoreLine()
		})

	labelRule := parser.NewRule().
		Pattern("(?P<indent>\\s*)(?P<label>[^\\s].+?)?(?:\\[(?P<link>.*)\\])?").
		Handler(func(ctx *parser.Context) {
			indent := ctx.Param("indent")
			label := strings.TrimSpace(ctx.Param("label"))
			link := strings.TrimSpace(ctx.Param("link"))

			if len(indent)%4 != 0 {
				ctx.Error("indentation must be in 4 space increments")
				return
			}
			indentLevel := len(indent) / 4

			err := tmp.AddItem(indentLevel, page.NewItem(link, label))
			if err != nil {
				ctx.Error(err.Error())
				return
			}
			ctx.IgnoreLine()
		})

	ps := parser.New()
	ps.Add(
		emptyRule,
		whitespaceRule,
		commentRule,
		versionRule,
		metadataRule,
		headerRule,
		groupRule,
		labelRule,
	)

	err := ps.Parse(spec)
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(tmp, "", "    ")
	fmt.Println(string(b))
}
