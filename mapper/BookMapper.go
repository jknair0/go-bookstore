package mapper

import (
	"tech.jknair/bookstore/db/schema"
	"tech.jknair/bookstore/model"
)

type BookMapper struct {
}

func CreateBookMapper() *BookMapper {
	return &BookMapper{}
}

func (b *BookMapper) ToData(domain ...*model.Book) []*schema.BookSchema {
	var books []*schema.BookSchema
	for _, item := range domain {
		books = append(books, &schema.BookSchema{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Author:    item.Author,
			CreatedAt: item.CreatedAt,
		})
	}
	return books
}

func (b *BookMapper) FromData(data ...*schema.BookSchema) model.Books {
	var books model.Books
	for _, item := range data {
		dbObj := item
		books = append(books, &model.Book{
			Uuid:      dbObj.Uuid,
			Name:      dbObj.Name,
			Author:    dbObj.Author,
			CreatedAt: dbObj.CreatedAt,
		})
	}
	return books
}
