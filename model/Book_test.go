package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBook_EncodeBook(t *testing.T) {
	name := "Deep work"
	author := "Carl Jung"
	book := CreateBook(name, author)
	bookJsonBytes, err := book.EncodeBook()
	assert.Nil(t, err)

	expectedJson := fmt.Sprintf(`{"uuid":"", "name":"%s", "author":"%s", "created_at":0}`, name, author)
	actualJson := string(bookJsonBytes)

	assert.JSONEq(t, expectedJson, actualJson)
}
