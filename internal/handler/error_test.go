package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestClientErr(t *testing.T) {
	t.Run("should have 400 status", func(t *testing.T) {
		err := ClientErr("error")
		if err.code != http.StatusBadRequest {
			t.Fatal("Error status is not 400")
		}
	})

	t.Run("should contain the error message", func(t *testing.T) {
		message := "error message"
		err := ClientErr(message)
		if err.Error() != message {
			t.Fatal("Error message does not match")
		}
	})
}

func TestInternalErr(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	t.Run("should have 500 status", func(t *testing.T) {
		err := InternalErr("error")
		if err.code != http.StatusInternalServerError {
			t.Fatal("Error status does not match")
		}
	})

	t.Run("should not contain the given error message", func(t *testing.T) {
		message := "error message"
		err := InternalErr(message)
		if err.Error() == message {
			t.Fatal("Error message should not match")
		}
	})
}
