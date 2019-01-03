package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
)

func TestHandler(t *testing.T) {
	t.Run("should not produce an error when given a valid page spec", func(t *testing.T) {
		res := &function.Response{}
		funcErr := handler(&function.Request{
			Body: "# test\nversion 1\n===\ntestlabel",
		}, res)
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}
	})

	t.Run("should respond with status code 400 if page spec is invalid", func(t *testing.T) {
		funcErr := handler(&function.Request{
			Body: "invalid spec",
		}, &function.Response{})
		if funcErr == nil {
			t.Fatal("Expected invalid spec to produce error")
		}

		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect status code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}
	})
}