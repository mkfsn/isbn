package isbn_test

import (
	"testing"

	"github.com/mkfsn/isbn"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidISBN(t *testing.T) {
	_, err := isbn.New("0")
	assert.NotNil(t, err)
}

func TestNewISBN10(t *testing.T) {
	book, err := isbn.New("0134190564")
	assert.Nil(t, err)
	assert.Equal(t, book.Info.Title, "The Go Programming Language")
}

func TestNewISBN13(t *testing.T) {
	book, err := isbn.New("9780134190563")
	assert.Nil(t, err)
	assert.Equal(t, book.Info.Title, "The Go Programming Language")
}
