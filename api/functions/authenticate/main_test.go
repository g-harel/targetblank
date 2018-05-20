package main

import (
	"fmt"
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/token"
)

func TestHandler(t *testing.T) {
	s := "test payload"
	tt, err := token.Seal([]byte(s))
	if err != nil {
		t.Fatalf("Unexpected error when sealing token: %v", err)
	}
	pp, err := token.Open(tt)
	if err != nil {
		t.Fatalf("Unexpected error when opening token: %v", err)
	}
	if s != string(pp) {
		t.Fatal("token does not match")
	}

	t.Run("", func(t *testing.T) {
		res := &function.Response{
			Headers: map[string]string{},
		}
		tok, err := function.MakeToken(false, "123456")
		if err != nil {
			t.Fatalf("Unexpected error when making token: %v", err)
		}
		fmt.Println(tok)
		handler(&function.Request{
			PathParameters: map[string]string{
				"address": "123456",
			},
			Headers: map[string]string{
				"Token": tok,
			},
			Body: "",
		}, res)
	})
}
