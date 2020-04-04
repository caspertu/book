package model

import "testing"

func TestCreate(t *testing.T) {
	expected := 2
	b := New()
	b.Create()
	actual := b.ID
	if expected != actual {
		t.Errorf("Create() actual return %d, expected %d", actual, expected)
	}
	t.Logf("book.Create() return %d", actual)
}

func TestFindBook(t *testing.T) {
	ID := 1
	actual := FindBook(ID)
	expected := New()
	if *expected != *actual {
		t.Errorf("Create() actual return %v, expected %v", actual, expected)
	}
	t.Logf("FindBook() return %v", actual)
}

func TestDelete(t *testing.T) {
	book := FindBook(1)
	actual := book.Delete()
	if !actual {
		t.Errorf("Delete() actual return %t", actual)
	}
	t.Logf("book.Delete() return %t", actual)
}

func TestUpdate(t *testing.T) {
	book := FindBook(1)
	actual := book.Update()
	if !actual {
		t.Errorf("Update() actual return %t", actual)
	}
	t.Logf("book.Update() return %t", actual)
}
