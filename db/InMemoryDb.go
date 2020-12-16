package db

type InMemoryDb struct {
	booksStore []BookSchema
}

func Create() InMemoryDb {
	return InMemoryDb{
		booksStore: []BookSchema{},
	}
}

func (i *InMemoryDb) SaveBooks(book []BookSchema) {
	i.booksStore = append(i.booksStore, book...)
}

func (i *InMemoryDb) GetBooks() []BookSchema {
	return i.booksStore
}

func (i *InMemoryDb) DeleteBook(uuid string) *BookSchema {
	deleteIndex := -1
	for index, book := range i.booksStore {
		if book.Uuid == uuid {
			deleteIndex = index
			break
		}
	}
	if deleteIndex == -1 {
		return nil
	}
	var deletedValue *BookSchema = nil
	if deleteIndex != -1 {
		deletedValue = &i.booksStore[deleteIndex]
		i.booksStore = append(i.booksStore[:deleteIndex], i.booksStore[deleteIndex+1:]...)
	}
	return deletedValue
}

func (i *InMemoryDb) UpdateBook(book BookSchema) *BookSchema {
	panic("implement me")
}
