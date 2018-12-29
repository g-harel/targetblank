package main

import (
	"fmt"
	"os"

	"github.com/g-harel/targetblank/internal/token"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("missing ciphertext")
	}

	msg, err := token.Open(args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(msg))
}
