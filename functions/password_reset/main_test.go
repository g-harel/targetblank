package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/function"
	mockMailer "github.com/g-harel/targetblank/services/mailer/mock"
	"github.com/g-harel/targetblank/services/storage"
	mockStorage "github.com/g-harel/targetblank/services/storage/mock"
)

func init() {
	mailerSend = mockMailer.Send
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

	t.Run("should check that the page's email matches", func(t *testing.T) {
		email := "oP8a0M2G@example.com"

		page := &storage.Page{
			Email: email,
		}
		_, err := mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		funcErr := handler(&function.Request{
			Body: "test@example.com",
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, &function.Response{})
		if funcErr == nil {
			t.Fatal("Expected handler to reject non-matching email")
		}
	})

	t.Run("should not change the page's password", func(t *testing.T) {
		email := "vKWA4GsS@example.com"

		h, err := crypto.Hash(email)
		if err != nil {
			t.Fatalf("Unexpected error when creating email hash: %v", err)
		}

		page := &storage.Page{
			Email: h,
		}
		_, err = mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		pass := page.Password

		funcErr := handler(&function.Request{
			Body: email,
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Handler failed: %v", funcErr)
		}

		page, err = mockStorage.PageRead(page.Addr)
		if err != nil {
			t.Fatalf("Unexpected error when fetching page: %v", err)
		}
		if page == nil {
			t.Fatal("Item does not exist")
		}

		if page.Password != pass {
			t.Fatalf("Item's password was changed \"%v\"", pass)
		}
	})

	t.Run("should send a confirmation email", func(t *testing.T) {
		email := "EDzhUzR8@example.com"

		h, err := crypto.Hash(email)
		if err != nil {
			t.Fatalf("Unexpected error when creating email hash: %v", err)
		}

		page := &storage.Page{
			Email: h,
		}
		_, err = mockStorage.PageCreate(page)
		if err != nil {
			t.Fatalf("Unexpected error when creating new page: %v", err)
		}

		funcErr := handler(&function.Request{
			Body: email,
			PathParameters: map[string]string{
				"addr": page.Addr,
			},
		}, &function.Response{})
		if funcErr != nil {
			t.Fatalf("Unexpected handler error: %v", funcErr)
		}

		e := mockMailer.LastSentTo(email)
		if e == nil {
			t.Fatal("No confirmation email was sent")
		}

		if strings.Index(e.Body, page.Addr) < 0 {
			t.Fatal("Confirmation email does not contain page's address")
		}
	})
}
