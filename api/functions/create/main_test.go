package main

import (
	"net/http"
	"testing"

	mockEmail "github.com/g-harel/targetblank/api/internal/email/mock"
	"github.com/g-harel/targetblank/api/internal/function"
	mockTables "github.com/g-harel/targetblank/api/internal/tables/mock"
)

func init() {
	pages = mockTables.NewPage()
	sender = mockEmail.NewSender()
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
