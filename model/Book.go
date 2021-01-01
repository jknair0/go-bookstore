package model

import (
	"bytes"
	"encoding/json"
	"strings"
)

type Book struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	CreatedAt int64  `json:"created_at"`
}

func CreateBook(name string, author string) *Book {
	return &Book{
		Uuid:      "",
		Name:      name,
		Author:    author,
		CreatedAt: 0,
	}
}

func EmptyBook() *Book {
	return CreateBook("", "")
}

func DecodeBook(b []byte) (*Book, error) {
	book := EmptyBook()
	err := json.NewDecoder(bytes.NewBuffer(b)).Decode(book)
	book.sanitizeValues()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *Book) sanitizeValues() {
	b.Name = strings.TrimSpace(b.Name)
	b.Author = strings.TrimSpace(b.Author)
}

func (b *Book) EncodeBook() ([]byte, error) {
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
