package mapper

import (
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/usecases/model"
)

type BookMapper struct {
}

func (b BookMapper) toData(data ...interface{}) []interface{} {
	var books []interface{}
	for _, item := range data {
		schema, ok := item.(db.BookSchema)
		if !ok {
			inputType := NewInvalidMapperInputType(item)
			panic(inputType.Error())
		}
		books = append(books, model.Book{
			Uuid:      schema.Uuid,
			Name:      schema.Name,
			Author:    schema.Author,
			CreatedAt: schema.CreatedAt,
		})
	}
	return books
}

func (b BookMapper) fromData(domain ...interface{}) []interface{} {
	var books []interface{}
	for _, item := range domain {
		schema, ok := item.(model.Book)
		if !ok {
			inputType := NewInvalidMapperInputType(item)
			panic(inputType.Error())
		}
		books = append(books, db.BookSchema{
			Uuid:      schema.Uuid,
			Name:      schema.Name,
			Author:    schema.Author,
			CreatedAt: schema.CreatedAt,
		})
	}
	return books
}
