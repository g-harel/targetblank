package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	storagePageDelete = mockStorage.PageDelete
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
				"token": "bad token",
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

	t.Run("should remove the page with the given address from the data store", func(t *testing.T) {
		page := &storage.Page{}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := function.CreateToken(false, page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		funcErr := handler(&function.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				"token": token,
			},
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		page, err = mockStorage.PageRead(page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when fetching page: %v", err)
		}
		if page != nil {
			t.Fatal("Item should not exist after being deleted")
		}
	})
}
