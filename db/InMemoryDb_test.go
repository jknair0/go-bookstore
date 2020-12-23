package db

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"tech.jknair/bookstore/db/schema"
	"tech.jknair/bookstore/testutils"
	"testing"
)

var RunTest func(func())

var subject *InMemoryDb

func setUp() {
	subject = CreateInMemoryDb()
}

func tearDown() {
	subject = nil
}

func TestMain(m *testing.M) {
	RunTest = testutils.CreateForEach(setUp, tearDown)
	m.Run()
	RunTest = nil
}

func TestInMemoryDb_SaveBooks(t *testing.T) {
	RunTest(func() {
		assert.Len(t, subject.GetBooks(), 0)
		if len(subject.GetBooks()) != 0 {
			t.Errorf("initial list is supposed to be empty")
			return
		}
		sampleBook1 := createSampleBook()
		uidArray := subject.SaveBooks([]*schema.BookSchema{&sampleBook1})
		assert.Len(t, uidArray, 1)
		booksList := subject.GetBooks()
		assert.Len(t, booksList, 1)
		actualResult := booksList[0]
		assert.Same(t, &sampleBook1, actualResult)
	})
}

func createSampleBook() schema.BookSchema {
	return schema.CreateBookSchema(uuid.New().String(), uuid.New().String())
}

func TestInMemoryDb_DeleteBook(t *testing.T) {
	RunTest(func() {
		sampleBook := createSampleBook()
		assert.Empty(t, subject.GetBooks())
		savedUuid := subject.SaveBooks([]*schema.BookSchema{&sampleBook})
		assert.Len(t, savedUuid, 1)
		assert.Len(t, subject.GetBooks(), 1)
		book := subject.DeleteBook(savedUuid[0])
		assert.NotNil(t, book)
		assert.Same(t, &sampleBook, book)
		assert.Empty(t, subject.GetBooks())
	})
}

func TestInMemoryDb_GetBook(t *testing.T) {
	RunTest(func() {
		sampleBook0 := createSampleBook()
		sampleBook1 := createSampleBook()
		sampleBook2 := createSampleBook()

		assert.Empty(t, subject.GetBooks())
		assert.Len(t, subject.SaveBooks([]*schema.BookSchema{&sampleBook0, &sampleBook1, &sampleBook2}), 3)

		t.Run("book exists", func(t *testing.T) {
			existingBooks := subject.GetBooks()
			aExistingBook := existingBooks[1]

			foundBook := subject.GetBook(aExistingBook.Uuid)
			assert.Equal(t, aExistingBook, foundBook)
		})

		t.Run("book not exists", func(t *testing.T) {
			foundBook := subject.GetBook(uuid.New().String())
			assert.Nil(t, foundBook)
		})
	})
}

func TestInMemoryDb_UpdateBook(t *testing.T) {
	RunTest(func() {
		sampleBook0 := createSampleBook()
		sampleBook1 := createSampleBook()
		updatedBook := createSampleBook()

		assert.Empty(t, subject.GetBooks())
		savedUuid := subject.SaveBooks([]*schema.BookSchema{&sampleBook0, &sampleBook1})
		assert.NotEmpty(t, savedUuid)
		assert.Len(t, subject.GetBooks(), 2)

		updated := subject.UpdateBook(&updatedBook)
		assert.True(t, updated)

		bookFound := false
		getBooksResult := subject.GetBooks()
		for _, book := range getBooksResult {
			bookFound = bookFound || cmp.Equal(book, &updatedBook)
		}
		assert.True(t, bookFound, "book not found in the updated list")
	})
}
