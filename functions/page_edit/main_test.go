package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/g-harel/targetblank/internal/handlers"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	storagePageUpdateDocument = mockStorage.PageUpdateDocument
}

func TestHandler(t *testing.T) {
	t.Run("should require an address param", func(t *testing.T) {
		err := handler(&handlers.Request{
			PathParameters: map[string]string{},
		}, &handlers.Response{})
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

	t.Run("should check that the page's email matches", func(t *testing.T) {
		email := "j8THwv6f@example.com"

		page := &storage.Page{
			Email: email,
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		funcErr := handler(&handlers.Request{
			Body: "test@example.com",
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, &handlers.Response{})
		if funcErr == nil {
			t.Fatal("Expected handler to reject non-matching email")
		}
	})

	t.Run("should change the page's page", func(t *testing.T) {
		label := "uMmETQtzy85kPOjU"

		page := &storage.Page{}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := handlers.CreateToken(false, page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		funcErr := handler(&handlers.Request{
			Body: "version 1\n===\n" + label,
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				"token": token,
			},
		}, &handlers.Response{})
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		page, err = mockStorage.PageRead(page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when fetching page: %v", err)
		}
		if page == nil {
			t.Fatal("Item does not exist")
		}

		if strings.Index(page.Document, label) < 0 {
			t.Fatal("Item's page was not changed")
		}
	})

	t.Run("should reject invalid page specs", func(t *testing.T) {
		page := &storage.Page{}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := handlers.CreateToken(false, page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		funcErr := handler(&handlers.Request{
			Body: "invalid spec",
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				"token": token,
			},
		}, &handlers.Response{})
		if funcErr == nil {
			t.Fatal("Expected invalid spec to produce an error")
		}
	})
}
