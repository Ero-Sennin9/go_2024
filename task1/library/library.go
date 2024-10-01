package library

import (
	"task1/library/book"
	"task1/library/storage"
)

type Library struct {
	storage.Storage
	Id book.Id
}

func (lib Library) Search(name string) (book.Book, bool) {
	return lib.Storage.Search(lib.Id(book.Book{Name: name}))
}

func (lib *Library) Add(book book.Book) {
	lib.Storage.Add(lib.Id(book), book)
}

func (lib *Library) SetId(id book.Id) {
	lib.Storage.ChangeId(id)
	lib.Id = id
}
