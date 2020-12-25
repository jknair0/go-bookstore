package mapper

import (
	"fmt"
	"reflect"
)

type InvalidMapperInputType struct {
	input interface{}
}

func NewInvalidMapperInputType(input interface{}) InvalidMapperInputType {
	return InvalidMapperInputType{input: input}
}

func (i *InvalidMapperInputType) Error() string {
	return fmt.Sprintf(
		"invalid input type. Expected db.BookSchema found %#v",
		reflect.TypeOf(i.input),
	)
}
