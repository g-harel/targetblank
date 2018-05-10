package main

import (
	"testing"

	"github.com/g-harel/targetblank/internal/function"
)

func TestHandler(t *testing.T) {
	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			PathParameters: map[string]string{
				"address": "123456",
			},
			Body: "test1234",
		}, res)
		// TODO use function config
		if res.StatusCode%100 != 2 {
			t.Fatalf(res.Body)
		}
	})
}
