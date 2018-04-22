package main

import (
	"encoding/json"
	"fmt"

	"github.com/g-harel/targetblank/page/parser"
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
	tmp, err := pageparser.Parse(spec)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(tmp, "", "    ")
	fmt.Println(string(b))
}
