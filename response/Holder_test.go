package response

import (
	"github.com/stretchr/testify/assert"
	"tech.jknair/bookstore/model"
	"testing"
)

func TestHolder_Encode(t *testing.T) {
	t.Run("data encoding", func(t *testing.T) {
		book := model.CreateBook("jk", "jk-author")
		responseHolder := CreateSuccessHolder(book)
		jsonBytes := responseHolder.EncodeJson()
		assert.JSONEq(t, `{"data":{"uuid":"","name":"jk","author":"jk-author","created_at":0},"error":null}`, string(jsonBytes))
	})
	t.Run("error encoding", func(t *testing.T) {
		responseHolder := CreateErrorHolder(ErrServerError)
		jsonBytes := responseHolder.EncodeJson()
		assert.JSONEq(t, `{"data":null,"error":{"error_code":3,"message":"Internal Server Error"}}`, string(jsonBytes))
	})
}
