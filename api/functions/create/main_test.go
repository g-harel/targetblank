package main

import (
	"net/http"
	"strings"
	"testing"

	mockEmail "github.com/g-harel/targetblank/api/internal/email/mock"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/rand"
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

	t.Run("should create a new page and respond with its key", func(t *testing.T) {
		email := rand.String(8) + "@example.com"

		res := &function.Response{}
		funcErr := handler(&function.Request{
			Body: email,
		}, res)
		if funcErr != nil {
			t.Fatalf("Unexpected handler error: %v", funcErr)
		}

		item, err := pages.Fetch(res.Body)
		if err != nil {
			t.Fatalf("Unexpected error when fetching new page item: %v", err)
		}
		if item == nil {
			t.Fatal("Item was not created")
		}

		ok := hash.Check(email, item.Email)
		if !ok {
			t.Fatal("Item's email hash does not match given one")
		}

		if item.Published {
			t.Fatal("New items should not be public by default")
		}
	})

	t.Run("should send a confirmation email", func(t *testing.T) {
		email := rand.String(8) + "@example.com"

		res := &function.Response{}
		funcErr := handler(&function.Request{
			Body: email,
		}, res)
		if funcErr != nil {
			t.Fatalf("Unexpected handler error: %v", funcErr)
		}

		e := mockEmail.LastSentTo(email)
		if e == nil {
			t.Fatal("No confirmation email was sent")
		}

		if strings.Index(e.Body, res.Body) < 0 {
			t.Fatal("Confirmation email does not contain new page's address")
		}
	})
}