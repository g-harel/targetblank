package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	mockSecrets "github.com/g-harel/targetblank/services/secrets/mock"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	secretsKey = mockSecrets.Key
	storagePageUpdatePassword = mockStorage.PageUpdatePassword
}

func TestPasswd(t *testing.T) {
	t.Run("should require an address param", func(t *testing.T) {
		err := Passwd(&handler.Request{
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

	t.Run("should require a validation token", func(t *testing.T) {
		err := Passwd(&handler.Request{
			PathParameters: map[string]string{
				"addr": "123456",
			},
			Headers: map[string]string{
				"token": "bad token",
			},
		}, &handler.Response{})
		if err == nil {
			t.Fatalf("Bad token should produce error")
		}
		if err.Code() != http.StatusNotFound {
			t.Fatalf(
				"Incorrect status code for bad token, expected %v but got %v: %v",
				http.StatusNotFound, err.Code(), err,
			)
		}

		err = Passwd(&handler.Request{
			PathParameters: map[string]string{
				"addr": "123456",
			},
			Headers: map[string]string{},
		}, &handler.Response{})
		if err == nil {
			t.Fatalf("Missing token should produce error")
		}
		if err.Code() != http.StatusNotFound {
			t.Fatalf(
				"Incorrect status code for missing token, expected %v but got %v: %v",
				http.StatusNotFound, err.Code(), err,
			)
		}
	})

	t.Run("should change the page's password", func(t *testing.T) {
		addr := "XMBofk"
		pass := "Uw9zJYVaSVxRkcac"

		page := &storage.Page{
			Addr: addr,
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		token, err := handler.CreateToken(mockSecrets.RawKey, false, addr)
		if err != nil {
			t.Fatalf("Unexpected error when creating token: %v", err)
		}

		funcErr := Passwd(&handler.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Headers: map[string]string{
				handler.AuthHeader: handler.AuthType + " " + token,
			},
			Body: pass,
		}, &handler.Response{})
		if funcErr != nil {
			t.Fatalf("Passwd failed: %v", funcErr)
		}

		page, err = mockStorage.PageRead(addr)
		if err != nil {
			t.Fatalf("Unexpected error when fetching page: %v", err)
		}
		if page == nil {
			t.Fatal("Item does not exist")
		}

		if !crypto.HashCheck(pass, page.Password) {
			t.Fatal("Fetched page's password does not match")
		}
	})

}
