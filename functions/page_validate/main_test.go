package main

import (
	"testing"

	"github.com/g-harel/targetblank/internal/function"
)

func TestHandler(t *testing.T) {
	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			Body: "version 1\n===\nexample.com",
		}, res)
	})
}
