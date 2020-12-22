package schema

type BookSchema struct {
	Uuid      string
	Name      string
	Author    string
	CreatedAt int64
}

func CreateBookSchema(name string, author string) BookSchema {
	return BookSchema{
		Uuid:      "",
		Name:      name,
		Author:    author,
		CreatedAt: 0,
	}
}
