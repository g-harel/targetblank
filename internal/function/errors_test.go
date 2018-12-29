package function

import (
	"errors"
	"net/http"
	"testing"
)

func TestCustomErr(t *testing.T) {
	t.Run("should have 400 status", func(t *testing.T) {
		err := CustomErr(errors.New("error"))
		if err.code != http.StatusBadRequest {
			t.Fatal("Error status is not 400")
		}
	})

	t.Run("should contain the error message", func(t *testing.T) {
		message := "error message"
		err := CustomErr(errors.New(message))
		if err.Error() != message {
			t.Fatal("Error message does not match")
		}
	})
}

func TestErr(t *testing.T) {
	t.Run("should have use the given status", func(t *testing.T) {
		status := http.StatusTeapot
		err := Err(status, errors.New("error"))
		if err.code != status {
			t.Fatal("Error status does not match")
		}
	})

	t.Run("should not contain the given error message", func(t *testing.T) {
		message := "error message"
		err := Err(0, errors.New(message))
		if err.Error() == message {
			t.Fatal("Error message should not match")
		}
	})
}
