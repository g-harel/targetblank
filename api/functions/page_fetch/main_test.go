package main

import (
	"fmt"
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
)

func TestHandler(t *testing.T) {
	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			PathParameters: map[string]string{
				"addr": "Nc6AFe",
			},
		}, res)
		fmt.Println(res)
	})
}
