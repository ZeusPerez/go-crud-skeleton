package errors

import "fmt"

type NotFound struct {
}

func (e NotFound) Error() string {
	// We are not displaying any further information for not giving clues to attackers.
	return "Resource not found"
}

type Internal struct {
	Err error
}

func (e Internal) Error() string {
	return fmt.Sprintf("Something bad happened: %s", e.Err)
}
