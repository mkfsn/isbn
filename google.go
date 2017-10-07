package isbn

import (
	"log"

	"github.com/dghubble/sling"
)

type QueryList struct {
	Kind       string      `json:"kind"`
	TotalItems int         `json:"totalItems"`
	Items      []QueryItem `json:"items"`
}

type ReadingModes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

type VolumnInfo struct {
	Title               string              `json:"title"`
	Authors             []string            `json:"authors"`
	PublishedDate       string              `json:"publishedDate"`
	Description         string              `json:"description"`
	PageCount           int                 `json:"pageCount"`
	IndustryIdentifiers IndustryIdentifiers `json:"industryIdentifiers"`
	ImageLinks          ImageLinks          `json:"imageLinks"`
	Language            string              `json:"language"`
	PrintType           string              `json:"printType"`
	MaturityRating      string              `json:"maturityRating"`
	AllowAnonLogging    bool                `json:"allowAnonLogging"`
	ContentVersion      string              `json:"contentVersion"`
	PreviewLink         string              `json:"previewLink"`
	InfoLink            string              `json:"infoLink"`
	CanonicalVolumeLink string              `json:"canonicalVolumeLink"`
	ReadingModes        ReadingModes        `json:"readingModes"`
}

type SaleInfo struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"saleability"`
}

type AccessInfo struct {
	Country                string `json:"country"`
	Viewability            string `json:"viewability"`
	Embeddable             bool   `json:"embeddable"`
	PublicDomain           bool   `json:"publicDomain"`
	TextToSpeechPermission string `json:"textToSpeechPermission"`
	WebReaderLink          string `json:"webReaderLink"`
	AccessViewStatus       string `json:"accessViewStatus"`
	QuoteSharingAllowed    bool   `json:"quoteSharingAllowed"`
	/*
	   "epub": {
	    "isAvailable": false
	   },
	   "pdf": {
	    "isAvailable": false
	   },
	*/
}

type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type IndustryIdentifiers []IndustryIdentifier

func (i IndustryIdentifiers) match(format Format, isbn string) bool {
	for _, identifier := range i {
		if identifier.Type == "ISBN_10" && format == ISBN10 {
			return isbn == identifier.Identifier
		}
		if identifier.Type == "ISBN_13" && format == ISBN13 {
			return isbn == identifier.Identifier
		}
	}
	return false
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type QueryItem struct {
	Kind       string     `json:"kind"`
	ID         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumnInfo VolumnInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo AccessInfo `json:"accessInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
}

func search(format Format, isbn string) (Book, error) {
	// https://www.googleapis.com/books/v1/volumes?q=4048685198
	var success QueryList
	_, err := sling.New().Get("https://www.googleapis.com/books/v1/volumes?q=" + isbn).ReceiveSuccess(&success)
	if err != nil {
		log.Println("Error:", err.Error())
		return Book{}, ErrNotFound
	}

	if success.TotalItems < 1 || len(success.Items) < 1 {
		return Book{}, ErrGoogleNotFound
	}

	for _, item := range success.Items {
		if item.VolumnInfo.IndustryIdentifiers.match(format, isbn) {
			return Book{item.VolumnInfo}, nil
		}
	}

	return Book{}, ErrGoogleNotFound
}

func SearchISBN10(isbn string) (Book, error) {
	return search(ISBN10, isbn)
}

func SearchISBN13(isbn string) (Book, error) {
	return search(ISBN13, isbn)
}
