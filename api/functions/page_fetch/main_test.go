package main

import (
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
)

func TestHandler(t *testing.T) {
	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			PathParameters: map[string]string{
				"address": "123456",
			},
		}, res)
	})
}
