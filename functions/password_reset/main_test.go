package main

import (
	"net/http"
	"strings"
	"testing"

	mockEmail "github.com/g-harel/targetblank/internal/email/mock"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/hash"
	"github.com/g-harel/targetblank/internal/rand"
	"github.com/g-harel/targetblank/internal/tables"
	mockTables "github.com/g-harel/targetblank/internal/tables/mock"
)

func init() {
	pages = mockTables.NewPage()
	sender = mockEmail.NewSender()
}

func TestHandler(t *testing.T) {
	t.Run("should require an address param", func(t *testing.T) {
		err := handler(&function.Request{
			PathParameters: map[string]string{},
		}, &function.Response{})
		if err == nil {
			t.Fatalf("Missing address produce error")
		}
		if err.Code() != http.StatusInternalServerError {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusInternalServerError, err.Code(), err,
			)
		}
	})

	t.Run("should check that the item's email matches", func(t *testing.T) {
		email := rand.String(8) + "@example.com"

		item := &tables.PageItem{
			Email: email,
		}
		err := pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		funcErr := handler(&function.Request{
			Body: "test@example.com",
			PathParameters: map[string]string{
				"addr": item.Key,
			},
		}, &function.Response{})
		if funcErr == nil {
			t.Fatal("Expected handler to reject non-matching email")
		}
	})

	t.Run("should change the item's password", func(t *testing.T) {
		email := rand.String(8) + "@example.com"

		h, err := hash.New(email)
		if err != nil {
			t.Fatalf("Unexpected error when creating email hash: %v", err)
		}

		item := &tables.PageItem{
			Email: h,
		}
		err = pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		pass := item.Password

		funcErr := handler(&function.Request{
			Body: email,
			PathParameters: map[string]string{
				"addr": item.Key,
			},
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		item, err = pages.Fetch(item.Key)
		if err != nil {
			t.Fatalf("Unexpected error when fetching item: %v", err)
		}
		if item == nil {
			t.Fatal("Item does not exist")
		}

		if item.Password == pass {
			t.Fatalf("Item's password was not changed \"%v\"", pass)
		}
	})

	t.Run("should send a confirmation email", func(t *testing.T) {
		email := rand.String(8) + "@example.com"

		h, err := hash.New(email)
		if err != nil {
			t.Fatalf("Unexpected error when creating email hash: %v", err)
		}

		item := &tables.PageItem{
			Email: h,
		}
		err = pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		funcErr := handler(&function.Request{
			Body: email,
			PathParameters: map[string]string{
				"addr": item.Key,
			},
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Unexpected handler error: %v", funcErr)
		}

		e := mockEmail.LastSentTo(email)
		if e == nil {
			t.Fatal("No confirmation email was sent")
		}

		if strings.Index(e.Body, item.Key) < 0 {
			t.Fatal("Confirmation email does not contain page's address")
		}
	})
}
