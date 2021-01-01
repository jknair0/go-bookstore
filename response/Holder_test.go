package response

import (
	"github.com/jknair0/bookstore/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHolder_Encode(t *testing.T) {
	t.Run("data encoding", func(t *testing.T) {
		book := model.CreateBook("jk", "jk-author")
		responseHolder := CreateFormattedSuccessResponse(book)
		jsonBytes := responseHolder.encodeJson()
		assert.JSONEq(t, `{"data":{"uuid":"","name":"jk","author":"jk-author","created_at":0},"error":null}`, string(jsonBytes))
	})
	t.Run("error encoding", func(t *testing.T) {
		responseHolder := NewErrorHolderFromErrCode(ErrServerError)
		jsonBytes := responseHolder.encodeJson()
		assert.JSONEq(t, `{"data":null,"error":{"error_code":3,"message":"Internal Server Error"}}`, string(jsonBytes))
	})
}
