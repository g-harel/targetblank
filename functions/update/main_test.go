package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/g-harel/targetblank/internal/handler"
	mockSecrets "github.com/g-harel/targetblank/services/secrets/mock"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	secretsKey = mockSecrets.Key
	storagePageUpdateDocument = mockStorage.PageUpdateDocument
}

func TestUpdate(t *testing.T) {
	t.Run("should require an address param", func(t *testing.T) {
		err := Update(&handler.Request{
			PathParameters: map[string]string{},
		}, &handler.Response{})
		if err == nil {
			t.Fatalf("Missing address produce error")
		}
		if err.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusBadRequest, err.Code(), err,
			)
		}
	})

	t.Run("should check that the page's email matches", func(t *testing.T) {
		email := "j8THwv6f@example.com"

		page := &storage.Page{
			Email: email,
		}
		err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		funcErr := Update(&handler.Request{
			Body: "test@example.com",
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, &handler.Response{})
		if funcErr == nil {
			t.Fatal("Expected handler to reject non-matching email")
		}
	})

	t.Run("should change the page's page", func(t *testing.T) {
		label := "uMmETQtzy85kPOjU"

		page := &storage.Page{}
		err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := handler.CreateToken(mockSecrets.RawKey, false, page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		funcErr := Update(&handler.Request{
			Body: "version 1\n===\n" + label,
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				handler.AuthHeader: handler.AuthType + " " + token,
			},
		}, &handler.Response{})
		if funcErr != nil {
			t.Fatalf("Update failed: %v", funcErr)
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

	t.Run("should reject invalid page document", func(t *testing.T) {
		page := &storage.Page{}
		err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := handler.CreateToken(mockSecrets.RawKey, false, page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		funcErr := Update(&handler.Request{
			Body: "invalid document",
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				handler.AuthHeader: handler.AuthType + " " + token,
			},
		}, &handler.Response{})
		if funcErr == nil {
			t.Fatal("Expected invalid document to produce an error")
		}
	})
}
