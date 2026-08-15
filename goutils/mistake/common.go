package mistake

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrEmptyConfig   = errors.New("config input is nil")
	ErrInValidConfig = errors.New("config input is invalid")
)

func NewErr(e error, s string) error {
	return fmt.Errorf("%s: %w", s, e)
}

// REST API errors -> StatusCode
func ErrToStatusCode(err error) int {
	return http.StatusInternalServerError
}
