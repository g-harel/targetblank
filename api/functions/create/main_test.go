package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/tables/mock"
)

func init() {
	pages = mock.NewPage()
}

func TestHandler(t *testing.T) {
	t.Run("should expect a valid email address in body", func(t *testing.T) {
		funcErr := handler(&function.Request{
			Body: "",
		}, &function.Response{})
		if funcErr == nil {
			t.Fatalf("Handler should reject empty email")
		}
		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}

		funcErr = handler(&function.Request{
			Body: "bad email address @example.com",
		}, &function.Response{})
		if funcErr == nil {
			t.Fatalf("Handler should reject empty email")
		}
		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}
	})

	t.Run("should create a new page with the specified email address", func(t *testing.T) {
		// TODO
	})

	t.Run("should send a confirmation email", func(t *testing.T) {
		// TODO
	})
}
