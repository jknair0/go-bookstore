package testutils

import (
	"github.com/stretchr/testify/mock"
	"strconv"
	"tech.jknair/bookstore/db/schema"
)

type MockDatabase struct {
	mock.Mock
}

func CreateMockDatabase() *MockDatabase {
	return new(MockDatabase)
}

func (m *MockDatabase) SaveBooks(books []*schema.BookSchema) []string {
	m.Called(books)
	var result []string
	for i := range books {
		result = append(result, "uuid-"+strconv.Itoa(i))
	}
	return result
}

func (m *MockDatabase) GetBooks() []*schema.BookSchema {
	panic("implement me")
}

func (m *MockDatabase) GetBook(uuid string) (*schema.BookSchema, error) {
	panic("implement me")
}

func (m *MockDatabase) DeleteBook(uuid string) (*schema.BookSchema, error) {
	panic("implement me")
}

func (m *MockDatabase) UpdateBook(book *schema.BookSchema) (*schema.BookSchema, error) {
	panic("implement me")
}

