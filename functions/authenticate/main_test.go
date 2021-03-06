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
	storagePageRead = mockStorage.PageRead
}

func TestAuthenticate(t *testing.T) {
	t.Run("should require an address param", func(t *testing.T) {
		err := Authenticate(&handler.Request{
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

	t.Run("should create a token for valid passwords", func(t *testing.T) {
		pass := "password123"
		h, err := crypto.Hash(pass)
		if err != nil {
			t.Fatal("Unexpected error when hashing password")
		}

		page := &storage.Page{
			Password: h,
		}
		err = mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		res := &handler.Response{}
		funcErr := Authenticate(&handler.Request{
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
			Body: pass,
		}, res)
		if funcErr != nil {
			t.Fatalf("Authenticate failed: %v", funcErr)
		}

		if res.Body == "" {
			t.Fatal("Response does not contain token")
		}
	})

	t.Run("should not create a token for invalid authentication", func(t *testing.T) {
		addr := "9k5Vhs"

		pass := "password123"
		h, err := crypto.Hash(pass)
		if err != nil {
			t.Fatal("Unexpected error when hashing password")
		}

		page := &storage.Page{
			Addr:     addr,
			Password: h,
		}
		err = mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		funcErr := Authenticate(&handler.Request{
			PathParameters: map[string]string{
				"addr": addr,
			},
			Body: "incorrect password",
		}, &handler.Response{})
		if funcErr == nil {
			t.Fatalf("Should produce an error if the password is invalid")
		}
		if funcErr.Code() != http.StatusNotFound {
			t.Fatalf(
				"Incorrect status code for password error, expected %v but got %v: %v",
				http.StatusNotFound, funcErr.Code(), funcErr,
			)
		}
	})
}
