package mapper

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"tech.jknair/bookstore/db/schema"
	"tech.jknair/bookstore/model"
	"testing"
)

func TestBookMapper_toData(t *testing.T) {
	t.Run("for non-empty list", func(t *testing.T) {
		sliceLen := 6

		var booksInput []*model.Book
		for i := 0; i < sliceLen; i++ {
			strInt := strconv.Itoa(i)
			booksInput = append(booksInput, &model.Book{
				Uuid:      "uuid" + strInt,
				Name:      "name" + strInt,
				Author:    "author" + strInt,
				CreatedAt: int64(i),
			})
		}

		result := CreateBookMapper().ToData(booksInput...)

		for i, actualResult := range result {
			strInt := strconv.Itoa(i)
			expectedResult := &schema.BookSchema{
				Uuid:      "uuid" + strInt,
				Name:      "name" + strInt,
				Author:    "author" + strInt,
				CreatedAt: int64(i),
			}
			assert.Equal(t, expectedResult, actualResult)
		}
	})
	t.Run("empty list", func(t *testing.T) {
		booksInput := make([]*model.Book, 0)
		result := CreateBookMapper().ToData(booksInput...)
		assert.NotNil(t, result)
	})
}

func TestBookMapper_fromData(t *testing.T) {
	sliceLen := 5

	var booksInput []*schema.BookSchema
	for i := 0; i < sliceLen; i++ {
		strInt := strconv.Itoa(i)
		booksInput = append(booksInput, &schema.BookSchema{
			Uuid:      "uuid" + strInt,
			Name:      "name" + strInt,
			Author:    "author" + strInt,
			CreatedAt: int64(i),
		})
	}

	result := CreateBookMapper().FromData(booksInput...)

	for i, actualResult := range result {
		strInt := strconv.Itoa(i)
		expectedResult := &model.Book{
			Uuid:      "uuid" + strInt,
			Name:      "name" + strInt,
			Author:    "author" + strInt,
			CreatedAt: int64(i),
		}
		assert.Equal(t, expectedResult, actualResult)
	}
}
