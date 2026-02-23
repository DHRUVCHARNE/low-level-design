package model

import "fmt"

type Book struct {
	title string
	author string
	isbn string
	isAvailable bool
}

func NewBook(title,author,isbn string) *Book {
	return &Book{title: title,author: author,isbn: isbn,isAvailable: true}
}

func (b *Book) BorrowBook() bool {
	b.isAvailable=false
	return false
}

func (b *Book) ReturnBook()  {
	b.isAvailable=true
}
 func (b *Book) DisplayInfo() {
	isAvailable:="Available"
	if(!b.isAvailable){
		isAvailable="Borrowed"
	}
	fmt.Printf("%s by %s (ISBN:%s) - %s\n",b.title,b.author,b.isbn,isAvailable)
 }
