package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/internal/handler"
)

func TestValidate(t *testing.T) {
	t.Run("should not produce an error when given a valid page document", func(t *testing.T) {
		res := &handler.Response{}
		funcErr := Validate(&handler.Request{
			Body: "# test\nversion 1\n===\ntestlabel",
		}, res)
		if funcErr != nil {
			t.Fatalf("Validate failed: %v", funcErr)
		}
	})

	t.Run("should respond with status code 400 if page document is invalid", func(t *testing.T) {
		funcErr := Validate(&handler.Request{
			Body: "invalid document",
		}, &handler.Response{})
		if funcErr == nil {
			t.Fatal("Expected invalid document to produce error")
		}

		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect status code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}
	})
}
