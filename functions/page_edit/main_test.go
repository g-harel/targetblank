package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	storagePageUpdateData = mockStorage.PageUpdateData
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
		email := "j8THwv6f@example.com"

		item := &storage.Page{
			Email: email,
		}
		_, err := mockStorage.PageCreate(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		funcErr := handler(&function.Request{
			Body: "test@example.com",
			PathParameters: map[string]string{
				"addr": item.Addr,
			},
		}, &function.Response{})
		if funcErr == nil {
			t.Fatal("Expected handler to reject non-matching email")
		}
	})

	t.Run("should change the item's page", func(t *testing.T) {
		label := "uMmETQtzy85kPOjU"

		item := &storage.Page{}
		_, err := mockStorage.PageCreate(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		token, funcErr := function.MakeToken(false, item.Addr)
		if funcErr != nil {
			t.Fatalf("Unexpected error when creating token: %v", funcErr)
		}

		funcErr = handler(&function.Request{
			Body: "version 1\n===\n" + label,
			PathParameters: map[string]string{
				"addr": item.Addr,
			},
			Headers: map[string]string{
				"token": token,
			},
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		item, err = mockStorage.PageRead(item.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when fetching item: %v", err)
		}
		if item == nil {
			t.Fatal("Item does not exist")
		}

		if strings.Index(item.Data, label) < 0 {
			t.Fatal("Item's page was not changed")
		}
	})

	t.Run("should reject invalid page specs", func(t *testing.T) {
		item := &storage.Page{}
		_, err := mockStorage.PageCreate(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		token, funcErr := function.MakeToken(false, item.Addr)
		if funcErr != nil {
			t.Fatalf("Unexpected error when creating token: %v", funcErr)
		}

		funcErr = handler(&function.Request{
			Body: "invalid spec",
			PathParameters: map[string]string{
				"addr": item.Addr,
			},
			Headers: map[string]string{
				"token": token,
			},
		}, &function.Response{})
		if funcErr == nil {
			t.Fatal("Expected invalid spec to produce an error")
		}
	})
}
