package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/internal/handler"
	mockSecrets "github.com/g-harel/targetblank/services/secrets/mock"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	secretsKey = mockSecrets.Key
	storagePageRead = mockStorage.PageRead
}

// TODO add test reading when password's last update is more recent than now.

func TestRead(t *testing.T) {
	t.Run("should require an address param", func(t *testing.T) {
		err := Read(&handler.Request{
			PathParameters: map[string]string{},
		}, &handler.Response{})
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

	t.Run("should fetch pages with the given address and token", func(t *testing.T) {
		doc := "test page"

		page := &storage.Page{
			Document:           doc,
			Published:          false,
			PasswordLastUpdate: "2006-01-02T15:04:05-0700",
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := handler.CreateToken(mockSecrets.RawKey, false, page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		res := &handler.Response{}
		funcErr := Read(&handler.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				handler.AuthHeader: handler.AuthType + " " + token,
			},
		}, res)
		if funcErr != nil {
			t.Fatalf("Read failed: %v", funcErr)
		}

		if res.Body != doc {
			t.Fatalf(
				"Incorrect page content, expected \"%v\" but got \"%v\"",
				page, res.Body,
			)
		}
	})

	t.Run("should fetch published pages without a token", func(t *testing.T) {
		doc := "test page"

		page := &storage.Page{
			Document:  doc,
			Published: true,
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		res := &handler.Response{}
		funcErr := Read(&handler.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, res)
		if funcErr != nil {
			t.Fatalf("Read failed: %v", funcErr)
		}

		if res.Body != doc {
			t.Fatalf(
				"Incorrect response content, expected \"%v\" but got \"%v\"",
				page, res.Body,
			)
		}
	})

	t.Run("should not fetch pages without a token", func(t *testing.T) {
		page := &storage.Page{
			Published: false,
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		res := &handler.Response{}
		funcErr := Read(&handler.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				"token": "bad token",
			},
		}, res)
		if funcErr == nil {
			t.Fatalf("Bad token should produce error")
		}
		if funcErr.Code() != http.StatusNotFound {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusNotFound, funcErr.Code(), funcErr,
			)
		}

		res = &handler.Response{}
		funcErr = Read(&handler.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{},
		}, res)
		if funcErr == nil {
			t.Fatalf("Missing token should produce error")
		}
		if funcErr.Code() != http.StatusNotFound {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusNotFound, funcErr.Code(), funcErr,
			)
		}
	})
}
