package model

import "fmt"

// Book 图书信息
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book{Title: %s, Author: %v}", b.Title, b.Author)
}
