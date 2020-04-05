package model

import (
	"errors"
	"fmt"
)

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
var 	maxID  int


func (b *Book) String() string {
	return fmt.Sprintf("Book{Title: %s, Author: %v, ID: %d}", b.Title, b.Author, b.ID)
}

func New() *Book {
	maxID += 1
	return &Book{Title: "Go", Author: "At", ID: maxID}
}

func (b *Book) Create() {
	if len(bookStore) == 0 {
		bookStore = make([]*Book, 0, 128)
	}
	bookStore = append(bookStore, b)
	fmt.Printf("Book Created, %v\n", b)
}

func FindAll() []*Book {
	return bookStore
}

func FindBook(ID int) (*Book, error) {
	for _, book := range bookStore {
		if book.ID == ID {
			return book, nil
		}
	}
	return nil, errors.New("not found")
}

func (b *Book) Delete() bool {
	for k, book := range bookStore {
		if b.ID == book.ID {
			bookStore = append(bookStore[:k], bookStore[k+1:]...)
			return true
		}
	}
	return false
}

func (b *Book) Update() bool {
	return true
}
