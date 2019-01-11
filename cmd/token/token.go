package main

import (
	"fmt"
	"os"

	"github.com/g-harel/targetblank/internal/handlers"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("missing token value")
	}

	t, err := handlers.CreateToken(false, args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
