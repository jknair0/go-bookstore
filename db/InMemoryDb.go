package db

import (
	"github.com/google/uuid"
	"tech.jknair/bookstore/db/schema"
	"time"
)

type InMemoryDb struct {
	booksStore []*schema.BookSchema
}

func CreateInMemoryDb() *InMemoryDb {
	return &InMemoryDb{
		booksStore: []*schema.BookSchema{},
	}
}

func (i *InMemoryDb) SaveBooks(books []*schema.BookSchema) []string {
	booksUid := make([]string, len(books))
	for index := range books {
		uid := uuid.New().String()
		books[index].Uuid = uid
		books[index].CreatedAt = time.Now().Unix()
		booksUid[index] = uid
	}
	i.booksStore = append(i.booksStore, books...)
	return booksUid
}

func (i *InMemoryDb) DeleteBook(uuid string) (*schema.BookSchema, error) {
	deleteIndex := -1
	for index, book := range i.booksStore {
		if book.Uuid == uuid {
			deleteIndex = index
			break
		}
	}
	var deletedValue *schema.BookSchema
	if deleteIndex == -1 {
		return deletedValue, nil
	}
	if deleteIndex != -1 {
		deletedValue = i.booksStore[deleteIndex]
		i.booksStore = append(i.booksStore[:deleteIndex], i.booksStore[deleteIndex+1:]...)
	}
	return deletedValue, nil
}

func (i *InMemoryDb) GetBooks() []*schema.BookSchema {
	return i.booksStore
}

func (i *InMemoryDb) GetBook(uuid string) (*schema.BookSchema, error) {
	panic("Not implemented")
}

func (i *InMemoryDb) UpdateBook(book *schema.BookSchema) (*schema.BookSchema, error) {
	panic("implement me")
}
