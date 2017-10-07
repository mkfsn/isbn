package isbn

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("Book not found")
	ErrGoogleNotFound = errors.New("Book not found via Google Books API")

	ErrInvalidISBNDigit = errors.New("Invalid ISBN digit")
)
