package db

import "tech.jknair/bookstore/db/schema"

type Database interface {

	SaveBooks(books []*schema.BookSchema) []string

	GetBooks() []*schema.BookSchema

	GetBook(uuid string) (*schema.BookSchema, error)

	DeleteBook(uuid string) (*schema.BookSchema, error)

	UpdateBook(book *schema.BookSchema) (*schema.BookSchema, error)

}
