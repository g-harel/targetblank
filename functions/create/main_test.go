package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handlers"
	mockMailer "github.com/g-harel/targetblank/services/mailer/mock"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	mailerSend = mockMailer.Send
	storagePageCreate = mockStorage.PageCreate
}

func TestHandler(t *testing.T) {
	t.Run("should expect a valid email address in body", func(t *testing.T) {
		funcErr := handler(&handlers.Request{
			Body: "",
		}, &handlers.Response{})
		if funcErr == nil {
			t.Fatalf("Handler should reject empty email")
		}
		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}

		funcErr = handler(&handlers.Request{
			Body: "bad email address @example.com",
		}, &handlers.Response{})
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

	t.Run("should create a new page and respond with its address", func(t *testing.T) {
		email := "s8yljnzo@example.com"

		res := &handlers.Response{}
		funcErr := handler(&handlers.Request{
			Body: email,
		}, res)
		if funcErr != nil {
			t.Fatalf("Unexpected handler error: %v", funcErr)
		}

		page, err := mockStorage.PageRead(res.Body)
		if err != nil {
			t.Fatalf("Unexpected error when fetching new page page: %v", err)
		}
		if page == nil {
			t.Fatal("Item was not created")
		}

		ok := crypto.HashCheck(email, page.Email)
		if !ok {
			t.Fatal("Item's email hash does not match given one")
		}

		if page.Published {
			t.Fatal("New pages should not be public by default")
		}
	})

	t.Run("should send a confirmation email", func(t *testing.T) {
		email := "QdJA8638@example.com"

		res := &handlers.Response{}
		funcErr := handler(&handlers.Request{
			Body: email,
		}, res)
		if funcErr != nil {
			t.Fatalf("Unexpected handler error: %v", funcErr)
		}

		e := mockMailer.LastSentTo(email)
		if e == nil {
			t.Fatal("No confirmation email was sent")
		}

		if strings.Index(e.Body, res.Body) < 0 {
			t.Fatal("Confirmation email does not contain new page's address")
		}
	})
}
