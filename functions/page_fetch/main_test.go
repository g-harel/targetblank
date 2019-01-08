package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	storagePageRead = mockStorage.PageRead
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

	t.Run("should fetch pages with the given address and token", func(t *testing.T) {
		doc := "test page"

		page := &storage.Page{
			Document:  doc,
			Published: false,
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, funcErr := function.MakeToken(false, page.Addr)
		if funcErr != nil {
			t.Fatalf("Unexpected error when creating token: %v", funcErr)
		}

		res := &function.Response{}
		funcErr = handler(&function.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				"token": token,
			},
		}, res)
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
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

		res := &function.Response{}
		funcErr := handler(&function.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, res)
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
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

		res := &function.Response{}
		funcErr := handler(&function.Request{
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
		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}

		res = &function.Response{}
		funcErr = handler(&function.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{},
		}, res)
		if funcErr == nil {
			t.Fatalf("Missing token should produce error")
		}
		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect error code, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}
	})
}
