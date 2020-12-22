package db

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"tech.jknair/bookstore/db/schema"
	"testing"
)

func TestInMemoryDb_SaveBooks(t *testing.T) {
	subject := CreateInMemoryDb()

	if len(subject.GetBooks()) != 0 {
		t.Errorf("initial list is supposed to be empty")
		return
	}

	input := schema.CreateBookSchema("Harry Potter", "JK Rowling")

	uidArray := subject.SaveBooks([]schema.BookSchema{input})

	if len(uidArray) != 1 {
		t.Errorf("empty uid returned after save expected uid of lenght 1")
		return
	}

	booksList := subject.GetBooks()
	booksLen := len(booksList)
	if booksLen != 1 {
		t.Errorf("expected books list of lenght 1 received %v", booksLen)
		return
	}

	actualResult := booksList[0]
	if !cmp.Equal(actualResult, input, cmpopts.IgnoreFields(schema.BookSchema{}, "Uuid")) {
		t.Errorf("\n expected %#v \n actual   %#v ", input, actualResult)
	}

}
