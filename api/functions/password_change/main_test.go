package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
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

	t.Run("should require a validation token", func(t *testing.T) {
		err := handler(&function.Request{
			PathParameters: map[string]string{
				"addr": "123456",
			},
			Headers: map[string]string{
				"Token": "bad token",
			},
		}, &function.Response{})
		if err == nil {
			t.Fatalf("Bad token should produce error")
		}
		if err.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect status code for bad token, expected %v but got %v: %v",
				http.StatusBadRequest, err.Code(), err,
			)
		}

		err = handler(&function.Request{
			PathParameters: map[string]string{
				"addr": "123456",
			},
			Headers: map[string]string{},
		}, &function.Response{})
		if err == nil {
			t.Fatalf("Missing token should produce error")
		}
		if err.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect status code for missing token, expected %v but got %v: %v",
				http.StatusBadRequest, err.Code(), err,
			)
		}
	})

	t.Run("should change the item's password", func(t *testing.T) {
		addr := rand.String(6)
		pass := rand.String(16)

		item := &tables.PageItem{
			Key: addr,
		}
		err := pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		token, funcErr := function.MakeToken(false, addr)
		if funcErr != nil {
			t.Fatalf("Unexpected error when creating token: %v", funcErr)
		}

		funcErr = handler(&function.Request{
			PathParameters: map[string]string{
				"addr": item.Key,
			},
			Headers: map[string]string{
				"Token": token,
			},
			Body: pass,
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		item, err = pages.Fetch(addr)
		if err != nil {
			t.Fatalf("Unexpected error when fetching item: %v", err)
		}
		if item == nil {
			t.Fatal("Item does not exist")
		}

		if !hash.Check(pass, item.Password) {
			t.Fatal("Fetched item's password does not match")
		}
	})

}