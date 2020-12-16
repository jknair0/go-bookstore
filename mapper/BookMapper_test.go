package mapper

import (
	"github.com/google/go-cmp/cmp"
	"strconv"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/usecases/model"
	"testing"
)

func TestBookMapper_fromData(t *testing.T) {
	sliceLen := 5

	var booksInput []interface{}
	for i := 0; i < sliceLen; i++ {
		strInt := strconv.Itoa(i)
		booksInput = append(booksInput, model.Book{
			Uuid:      "uuid" + strInt,
			Name:      "name" + strInt,
			Author:    "author" + strInt,
			CreatedAt: int64(i),
		})
	}

	result := BookMapper{}.fromData(booksInput...)

	for i, actualResult := range result {
		strInt := strconv.Itoa(i)
		expectedResult := db.BookSchema{
			Uuid:      "uuid" + strInt,
			Name:      "name" + strInt,
			Author:    "author" + strInt,
			CreatedAt: int64(i),
		}
		assertion := cmp.Equal(actualResult, expectedResult)
		if !assertion {
			t.Errorf("assertion failed %#v != %#v", actualResult, expectedResult)
		}
	}
}

func TestBookMapper_toData(t *testing.T) {
	sliceLen := 6

	var booksInput []interface{}
	for i := 0; i < sliceLen; i++ {
		strInt := strconv.Itoa(i)
		booksInput = append(booksInput, db.BookSchema{
			Uuid:      "uuid" + strInt,
			Name:      "name" + strInt,
			Author:    "author" + strInt,
			CreatedAt: int64(i),
		})
	}

	result := BookMapper{}.toData(booksInput...)

	for i, actualResult := range result {
		strInt := strconv.Itoa(i)
		expectedResult := model.Book{
			Uuid:      "uuid" + strInt,
			Name:      "name" + strInt,
			Author:    "author" + strInt,
			CreatedAt: int64(i),
		}
		assertion := cmp.Equal(actualResult, expectedResult)
		if !assertion {
			t.Errorf("assertion failed %#v != %#v", actualResult, expectedResult)
		}
	}
}
