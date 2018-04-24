package parser

import (
	"fmt"
)

// Error adds line number to parser errors.
type Error struct {
	line  int
	cause string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error on line %d: %v", e.line, e.cause)
}
