package model

import "fmt"

// Book 图书信息
type Book struct {
	Title  string
	Author string
}

var bookStore []*Book

func (b *Book) String() string {
	return fmt.Sprintf("Book{Title: %s, Author: %v}", b.Title, b.Author)
}

func (b *Book) Create() (ID int) {
	if len(bookStore) == 0 {
		bookStore = make([]*Book, 0, 128)
		bookStore = append(bookStore, b)
	} else {
		bookStore = append(bookStore, b)
	}
	return len(bookStore)
}

func FindBook(ID int) *Book {
	//return Book{"Go", "At"}
	return bookStore[ID - 1]
}

func (b *Book) Delete() bool {
	return true
}

func (b *Book) Update() bool {
	return true
}
