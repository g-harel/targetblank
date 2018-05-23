package main

import (
	"net/http"
	"testing"

	"github.com/g-harel/targetblank/api/internal/tables"

	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/tables/mock"
)

func init() {
	pages = mock.NewPage()
}

func TestHandler(t *testing.T) {
	t.Run("It should remove the item with the given address from the table", func(t *testing.T) {
		item := &tables.PageItem{}
		err := pages.Create(item)
		if err != nil {
			t.Fatalf("Unexpected error when creating new item: %v", err)
		}

		item, err = pages.Fetch(item.Key)
		if err != nil {
			t.Fatalf("Unexpected error when fetching item: %v", err)
		}
		if item == nil {
			t.Fatal("Fetched item is empty")
		}

		token, funcErr := function.MakeToken(false, item.Key)
		if funcErr != nil {
			t.Fatalf("Unexpected function error when creating token: %v", funcErr)
		}

		res := &function.Response{
			StatusCode: http.StatusOK,
		}
		handler(&function.Request{
			PathParameters: map[string]string{
				"addr": item.Key,
			},
			Headers: map[string]string{
				"Token": token,
			},
		}, res)

		if res.StatusCode != http.StatusOK {
			t.Fatalf("Delete operation failed, incorrect status code %v", res.StatusCode)
		}

		item, err = pages.Fetch(item.Key)
		if err != nil {
			t.Fatalf("Unexpected error when fetching item: %v", err)
		}
		if item != nil {
			t.Fatal("Item should not exist after being deleted")
		}
	})
}
