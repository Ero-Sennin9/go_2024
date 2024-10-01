package map_storage

import "task1/library/book"

type MapStorage struct {
	data map[string]book.Book
}

func (mapStorage MapStorage) Search(id string) (book.Book, bool) {
	book, result := mapStorage.data[id]
	return book, result
}

func (mapStorage *MapStorage) Add(id string, bookItem book.Book) {
	if mapStorage.data == nil {
		mapStorage.data = make(map[string]book.Book)
	}
	mapStorage.data[id] = bookItem
}

func (mapStorage *MapStorage) ChangeId(id book.Id) {
	newData := make(map[string]book.Book)
	for _, book := range mapStorage.data {
		newData[id(book)] = book
	}
	mapStorage.data = newData
}

func MakeMapStorage() *MapStorage {
	pointer := new(MapStorage)
	pointer.data = make(map[string]book.Book)
	return pointer
}
