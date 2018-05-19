package parser

import (
	"fmt"
)

// Error adds line number to parser errors.
type Error struct {
	error
	Line int
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error on line %v: %v", e.Line, e.error)
}
