package db

import "github.com/jknair0/bookstore/db/schema"

type Database interface {
	SaveBooks(books []*schema.BookSchema) []string

	GetBooks() []*schema.BookSchema

	GetBook(uuid string) *schema.BookSchema

	DeleteBook(uuid string) *schema.BookSchema

	UpdateBook(book *schema.BookSchema) bool
}
