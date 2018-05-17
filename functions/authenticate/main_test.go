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
		t.Fatalf(err.Error())
	}
	pp, err := token.Open(tt)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if s != string(pp) {
		t.Fatalf("token does not match")
	}

	t.Run("", func(t *testing.T) {
		res := &function.Response{
			Headers: map[string]string{},
		}
		tok, err := function.MakeToken(false, "123456")
		if err != nil {
			t.Fatalf(err.Error())
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
