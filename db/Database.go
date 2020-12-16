package db

type BookSchema struct {
	Uuid string
	Name string
	Author string
	CreatedAt int64
}

type Database interface {

	SaveBooks(book []BookSchema)

	GetBooks() []BookSchema

	DeleteBook(uuid string) *BookSchema

	UpdateBook(book BookSchema) *BookSchema

}
