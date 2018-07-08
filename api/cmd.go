package main

import (
	"fmt"
	"os"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/token"
)

func fatal(args ...interface{}) {
	fmt.Println(args)
	os.Exit(1)
}

func makeToken(addr string) {
	t, err := function.MakeToken(false, addr)
	if err != nil {
		fatal("error", err)
	}
	fmt.Println(t)
}

func decryptError(e string) {
	msg, err := token.Open(e)
	if err != nil {
		fatal("error", err)
	}
	fmt.Println(string(msg))
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fatal("missing command")
	}

	command := args[0]
	args = args[1:]

	switch command {
	case "token":
		if len(args) < 1 {
			fatal("missing address")
		}
		makeToken(args[0])
	case "error":
		if len(args) < 1 {
			fatal("missing error string")
		}
		decryptError(args[0])
	}
}
