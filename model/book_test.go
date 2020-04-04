package model

import "testing"

func TestCreate(t *testing.T) {
	expected := 1
	b := Book{"Go", "At"}
	actual := b.Create()
	if expected != actual {
		t.Errorf("Create() actual return %d, expected %d", actual, expected)
	}
	t.Logf("book.Create() return %d", actual)
}

func TestFindBook(t *testing.T) {
	book := Book{"Go", "At"}
	ID := book.Create()
	actual := FindBook(ID)
	expected := &Book{"Go", "At"}
	if *expected != *actual {
		t.Errorf("Create() actual return %v, expected %v", actual, expected)
	}
	t.Logf("FindBook() return %v", actual)
}

func TestDelete(t *testing.T) {
	expected := true
	book := FindBook(1)
	actual := book.Delete()
	if expected != actual {
		t.Errorf("Delete() actual return %t, expected %t", actual, expected)
	}
	t.Logf("book.Delete() return %t", actual)
}

func TestUpdate(t *testing.T) {
	expected := true
	book := FindBook(1)
	actual := book.Update()
	if expected != actual {
		t.Errorf("Update() actual return %t, expected %t", actual, expected)
	}
	t.Logf("book.Update() return %t", actual)
}
