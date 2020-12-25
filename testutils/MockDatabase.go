package testutils

import (
	"github.com/jknair0/bookstore/db/schema"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func CreateMockDatabase() *MockDatabase {
	return new(MockDatabase)
}

func (m *MockDatabase) SaveBooks(books []*schema.BookSchema) []string {
	args := m.Called(books)
	return args.Get(0).([]string)
}

func (m *MockDatabase) GetBooks() []*schema.BookSchema {
	args := m.Called()
	return args.Get(0).([]*schema.BookSchema)
}

func (m *MockDatabase) GetBook(uuid string) *schema.BookSchema {
	args := m.Called(uuid)
	book := args.Get(0)
	if book == nil {
		return nil
	}
	return book.(*schema.BookSchema)
}

func (m *MockDatabase) DeleteBook(uuid string) *schema.BookSchema {
	panic("implement me")
}

func (m *MockDatabase) UpdateBook(book *schema.BookSchema) bool {
	panic("implement me")
}
