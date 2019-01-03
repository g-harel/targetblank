package main

import (
	"fmt"
	"os"

	"github.com/g-harel/targetblank/internal/crypto"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("missing ciphertext")
	}

	msg, err := crypto.Decrypt(args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(msg))
}
