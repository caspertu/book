package model

import "testing"

var bookOne = Book{"Go", "At", 1}

func TestCreate(t *testing.T) {
	expected := 2
	b := New()
	b.Create()
	actual := b.ID
	if expected != actual {
		t.Errorf("Create() actual return %d, expected %d", actual, expected)
	}
	t.Logf("book.Create(), book.ID = %d", actual)
}

func TestFindBook(t *testing.T) {
	ID := 1
	actual, _ := FindBook(ID)
	expected := &bookOne
	if *expected != *actual {
		t.Errorf("FindBook() actual return %v, expected %v", actual, expected)
	}
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