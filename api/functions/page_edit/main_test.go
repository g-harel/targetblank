package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/rand"
	"github.com/g-harel/targetblank/api/internal/tables"
	mockTables "github.com/g-harel/targetblank/api/internal/tables/mock"
)

func init() {
	pages = mockTables.NewPage()
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

	t.Run("should change the item's page", func(t *testing.T) {
		label := rand.String(16)

		item := &tables.PageItem{}
		err := pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		token, funcErr := function.MakeToken(false, item.Key)
		if funcErr != nil {
			t.Fatalf("Unexpected error when creating token: %v", funcErr)
		}

		funcErr = handler(&function.Request{
			Body: "version 1\n===\n" + label,
			PathParameters: map[string]string{
				"addr": item.Key,
			},
			Headers: map[string]string{
				"Token": token,
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

		if strings.Index(item.Page, label) < 0 {
			t.Fatal("Item's page was not changed")
		}
	})

	t.Run("should reject invalid page specs", func(t *testing.T) {
		item := &tables.PageItem{}
		err := pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		token, funcErr := function.MakeToken(false, item.Key)
		if funcErr != nil {
			t.Fatalf("Unexpected error when creating token: %v", funcErr)
		}

		funcErr = handler(&function.Request{
			Body: "invalid spec",
			PathParameters: map[string]string{
				"addr": item.Key,
			},
			Headers: map[string]string{
				"Token": token,
			},
		}, &function.Response{})
		if funcErr == nil {
			t.Fatal("Expected invalid spec to produce an error")
		}
	})
}
