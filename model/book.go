package model

import "fmt"

func init() {
	New().Create()
}

// Book 图书信息
type Book struct {
	Title  string
	Author string
	ID     int
}

var bookStore []*Book

func (b *Book) String() string {
	return fmt.Sprintf("Book{Title: %s, Author: %v, ID: %d}", b.Title, b.Author, b.ID)
}

func New() *Book {
	return &Book{Title: "Go", Author: "At", ID: 1}
}

func (b *Book) Create() {
	if len(bookStore) == 0 {
		bookStore = make([]*Book, 0, 128)
	}
	bookStore = append(bookStore, b)
	b.ID = len(bookStore)
	fmt.Printf("Book Created, ID=%d", b.ID)
}

func FindAll() []*Book {
	return bookStore
}

func FindBook(ID int) *Book {
	//return Book{"Go", "At"}
	return bookStore[ID-1]
}

func (b *Book) Delete() bool {
	return true
}

func (b *Book) Update() bool {
	return true
}
