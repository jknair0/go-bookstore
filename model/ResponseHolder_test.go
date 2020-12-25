package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponseHolder_EncodeResponseHolder(t *testing.T) {
	t.Run("data encoding", func(t *testing.T) {
		book := CreateBook("jk", "jk-author")
		responseHolder := CreateSuccessResponseHolder(book)
		jsonBytes := responseHolder.EncodeJson()
		assert.JSONEq(t, `{"data":{"uuid":"","name":"jk","author":"jk-author","created_at":0},"error":null}`, string(jsonBytes))
	})
	t.Run("error encoding", func(t *testing.T) {
		responseHolder := CreateErrorResponseHolder("An Error string")
		jsonBytes := responseHolder.EncodeJson()
		assert.JSONEq(t, `{"data":null,"error":"An Error string"}`, string(jsonBytes))
	})
}
