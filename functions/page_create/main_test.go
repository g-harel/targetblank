package main

import (
	"fmt"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
)

func TestHandler(t *testing.T) {
	// TODO mock database package
	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			Body: "test@gmail.com",
		}, res)
		fmt.Println(res)
	})
}
