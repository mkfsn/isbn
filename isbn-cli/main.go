package main

import (
	"fmt"
	"os"

	"github.com/mkfsn/isbn"
)

func usage() {
	fmt.Println("./isbn-cli [ISBN]")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 1 {
		usage()
	}

	isbnStr := os.Args[1]
	book, err := isbn.New(isbnStr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("%+v\n", book.Info)
}
