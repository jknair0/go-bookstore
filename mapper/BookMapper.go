package mapper

import (
	"github.com/jknair0/bookstore/db/schema"
	"github.com/jknair0/bookstore/model"
)

type BookMapper struct {
}

func CreateBookMapper() *BookMapper {
	return &BookMapper{}
}

func (b *BookMapper) ToData(domain ...*model.Book) []*schema.BookSchema {
	books := make([]*schema.BookSchema, len(domain))
	for index, item := range domain {
		books[index] = &schema.BookSchema{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Author:    item.Author,
			CreatedAt: item.CreatedAt,
		}
	}
	return books
}

func (b *BookMapper) FromData(data ...*schema.BookSchema) []*model.Book {
	books := make([]*model.Book, len(data))
	for index, item := range data {
		books[index] = &model.Book{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Author:    item.Author,
			CreatedAt: item.CreatedAt,
		}
	}
	return books
}
