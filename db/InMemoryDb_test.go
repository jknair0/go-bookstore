package db

import (
	"github.com/stretchr/testify/assert"
	"tech.jknair/bookstore/db/schema"
	"testing"
)

var subject *InMemoryDb

func TestMain(m *testing.M) {
	subject = CreateInMemoryDb()
	m.Run()
	subject = nil
}

func TestInMemoryDb_SaveBooks(t *testing.T) {
	assert.Len(t, subject.GetBooks(), 0)
	if len(subject.GetBooks()) != 0 {
		t.Errorf("initial list is supposed to be empty")
		return
	}
	sampleBook1 := sampleBook()
	uidArray := subject.SaveBooks([]*schema.BookSchema{&sampleBook1})
	assert.Len(t, uidArray, 1)
	booksList := subject.GetBooks()
	assert.Len(t, booksList, 1)
	actualResult := booksList[0]
	assert.Same(t, &sampleBook1, actualResult)
}

func sampleBook() schema.BookSchema {
	return schema.CreateBookSchema("Harry Potter", "JK Rowling")
}

func TestInMemoryDb_DeleteBook(t *testing.T) {
	sampleBook := sampleBook()
	assert.Empty(t, subject.GetBooks())
	savedUuid := subject.SaveBooks([]*schema.BookSchema{ &sampleBook })
	assert.Len(t, savedUuid, 1)
	assert.Len(t, subject.GetBooks(), 1)
	book, err := subject.DeleteBook(savedUuid[0])
	assert.Nil(t, err)
	assert.NotNil(t, book)
	assert.Same(t, &sampleBook, book)
	assert.Empty(t, subject.GetBooks())
}
