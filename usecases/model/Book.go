package model

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	CreatedAt int64  `json:"created_at"`
}

func CreateBook(name string, author string) Book {
	return Book{
		Uuid:      uuid.New().String(),
		Name:      name,
		Author:    author,
		CreatedAt: time.Now().Unix(),
	}
}

func EmptyBook() Book {
	return CreateBook("", "")
}