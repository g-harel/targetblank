package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/rand"
	"github.com/g-harel/targetblank/api/internal/tables"
	"github.com/g-harel/targetblank/api/internal/tables/mock"
)

func init() {
	pages = mock.NewPage()
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

	t.Run("should create a token for valid passwords", func(t *testing.T) {
		p := "password123"
		h, err := hash.New(p)
		if err != nil {
			t.Fatal("Unexpected error when hashing password")
		}

		item := &tables.PageItem{
			Password: h,
		}
		err = pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		res := &function.Response{}
		funcErr := handler(&function.Request{
			PathParameters: map[string]string{
				"addr": item.Key,
			},
			Body: p,
		}, res)
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		if res.Body == "" {
			t.Fatal("Response does not contain token")
		}
	})

	t.Run("should not create a token for invalid authentication", func(t *testing.T) {
		addr := rand.String(6)

		p := "password123"
		h, err := hash.New(p)
		if err != nil {
			t.Fatal("Unexpected error when hashing password")
		}

		item := &tables.PageItem{
			Key:      addr,
			Password: h,
		}
		err = pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		funcErr := handler(&function.Request{
			PathParameters: map[string]string{
				"addr": addr,
			},
			Body: "bad password",
		}, &function.Response{})
		if funcErr == nil {
			t.Fatalf("Should produce an error if the password is invalid")
		}
		if funcErr.Code() != http.StatusBadRequest {
			t.Fatalf(
				"Incorrect status code for password error, expected %v but got %v: %v",
				http.StatusBadRequest, funcErr.Code(), funcErr,
			)
		}
	})
}
