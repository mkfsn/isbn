package isbn

type Book struct {
	Info VolumnInfo
}

func New(isbn string) (Book, error) {
	// https://www.googleapis.com/books/v1/volumes?q=4048685198
	switch Identify(isbn) {
	case ISBN10:
		return SearchISBN10(isbn)
	case ISBN13:
		return SearchISBN13(isbn)
	}

	return Book{}, ErrNotFound
}
