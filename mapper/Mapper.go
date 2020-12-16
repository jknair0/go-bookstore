package mapper

type Mapper interface {
	toData(data ...interface{}) []interface{}

	fromData(domain ...interface{}) []interface{}
}

