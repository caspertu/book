package model

import (
	"errors"
	"fmt"
	"github.com/caspertu/book/douban"
)

func init() {
	New().Create()
}

// Book 图书信息
type Book struct {
	ID      int    `json:"id"`
	Barcode string `json:"barcode"`
	Title   string `json:"title"`
	Author  string `json:"author"`
}

var bookStore []*Book
var maxID int

func (b *Book) String() string {
	return fmt.Sprintf("Book{ID: %d, Barcode: %s, Title: %s, Author: %v}",
		b.ID, b.Barcode, b.Title, b.Author)
}

func New() *Book {
	book := &Book{ID: maxID, Barcode: "", Title: "Go", Author: "At"}
	book.GenerateID()
	return book
}

// NewBookWithBarcode ...
func NewBookWithBarcode(barcode string) *Book {
	book := New()
	book.Barcode = barcode
	return book
}

func (b *Book) GenerateID() int {
	if b.ID == 0 {
		maxID += 1
		b.ID = maxID
	}
	return maxID
}

func (b *Book) Scan() *Book {
	b.Barcode = "9787115431356"
	return b
}

func (b *Book) GetInfo() *Book {
	doubanBook := douban.New(b.Barcode)
	b.Title = doubanBook.Title
	var author string
	for _, a := range doubanBook.Author {
		author = author + a
	}
	b.Author = author
	return b
}

func (b *Book) Create() {
	b.GenerateID()
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
