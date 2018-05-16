package main

import (
	"fmt"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/token"
)

func TestHandler(t *testing.T) {
	s := "test payload"
	tt, err := token.Seal([]byte(s))
	if err != nil {
		t.Fatalf("!! %v", err)
	}
	fmt.Println(tt)
	pp, err := token.Open(tt)
	if err != nil {
		t.Fatalf("!! %v", err)
	}
	if s != string(pp) {
		t.Fatalf("didn't work")
	}

	t.Run("", func(t *testing.T) {
		res := &function.Response{}
		handler(&function.Request{
			PathParameters: map[string]string{
				"address": "123456",
			},
			Body: "test123",
		}, res)
	})
}
