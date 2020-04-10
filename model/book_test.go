package model

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

var bookOne = Book{ID: 1, Barcode: "", Title: "Go", Author: "At"}

func TestCreate(t *testing.T) {
	expected := 1
	b := New()
	b.Create()
	b.Barcode = "123"
	actual := b.ID
	assert.Equal(t, actual, expected)
	//t.Logf("book.Create(), book.ID = %d", actual)
}

func TestFindBook(t *testing.T) {
	ID := 1
	actual, _ := FindBook(ID)
	expected := &bookOne
	assert.Equal(t, actual, expected)
	t.Logf("FindBook() return %v", actual)
}

func TestUpdate(t *testing.T) {
	ID := 1
	book, err := FindBook(ID)
	if err != nil {
		t.Fatalf("FindBook(%d) not found", ID)
	}
	actual := book.Update()
	if !actual {
		t.Errorf("Update() actual return %t", actual)
	}
	t.Logf("book.Update() return %t", actual)
}

func TestDelete(t *testing.T) {
	book, _ := FindBook(1)
	actual := book.Delete()
	if !actual {
		t.Errorf("Delete() actual return %t", actual)
	}
	t.Logf("book.Delete() return %t", actual)
}

func TestNewWithBarcode(t *testing.T) {
	barcode := "9787115431356"
	book := NewBookWithBarcode(barcode)
	t.Log(book.Barcode)
	book = book.GetInfo()
	t.Log(book.Title)
}
