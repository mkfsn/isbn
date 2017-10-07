package isbn_test

import (
	"testing"

	"github.com/mkfsn/isbn"

	"github.com/stretchr/testify/assert"
)

func TestValidISBN10(t *testing.T) {
	cases := []string{
		"4048685198",
		"4167819058",
		"416781904X",
		"4167819031",
		"4167819023",
		"4840236968",
	}

	for _, c := range cases {
		assert.True(t, isbn.Is10(c))
	}
}

func TestInvalidISBN10(t *testing.T) {
	cases := []string{
		"4048685199",
		"4167819059",
		"4167819040",
		"4167819032",
		"4167819024",
		"4840236969",
		"ABCDEFGHIJ",
	}

	for _, c := range cases {
		assert.False(t, isbn.Is10(c))
	}
}

func TestValidISBN13(t *testing.T) {
	cases := []string{
		"978-4048685191",
		"978-4167819057",
		"978-4167819040",
		"978-4167819033",
		"978-4167819026",
		"978-4840236966",
	}

	for _, c := range cases {
		assert.True(t, isbn.Is13(c))
	}
}

func TestInvalidISBN13(t *testing.T) {
	cases := []string{
		"978-4048685192",
		"978-4167819058",
		"978-4167819041",
		"978-4167819034",
		"978-4167819027",
		"978-4840236967",
		"ABC-DEFGHIJKLM",
	}

	for _, c := range cases {
		assert.False(t, isbn.Is13(c))
	}
}

func BenchmarkIsISBN10(b *testing.B) {
	cases := []string{
		"4048685198",
		"4167819058",
		"416781904X",
		"4167819031",
		"4167819023",
		"4840236968",
	}

	for i := 0; i < b.N; i++ {
		isbn.Is10(cases[i%len(cases)])
	}
}

func BenchmarkIsISBN13(b *testing.B) {
	cases := []string{
		"978-4048685191",
		"978-4167819057",
		"978-4167819040",
		"978-4167819033",
		"978-4167819026",
		"978-4840236966",
	}

	for i := 0; i < b.N; i++ {
		isbn.Is13(cases[i%len(cases)])
	}
}
