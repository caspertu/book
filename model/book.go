package model

import "fmt"

// Book 图书信息
type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf("Book{Title: %s, Author: %v}", b.Title, b.Author)
}

func (b *Book) Create() (ID int) {
	ID = 1
	return
}

func FindBook() Book {
	return Book{"Go", "At"}
}

func (b *Book) Delete() bool {
	return true
}

func (b *Book) Update() bool {
	return true
}
