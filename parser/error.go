package parser

import (
	"fmt"
)

// Error adds line number to parser errors.
type Error struct {
	Line  int
	cause string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error on line %v: %v", e.Line, e.cause)
}
