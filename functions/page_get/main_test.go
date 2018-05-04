package main

import (
	"fmt"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
)

func TestHandler(t *testing.T) {
	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			Body: "123456",
		}, res)
		fmt.Println(res)
	})
}
