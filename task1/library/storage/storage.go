package storage

import "task1/library/book"

type Storage interface {
	Search(id string) (book.Book, bool)
	Add(id string, book book.Book)
	ChangeId(id book.Id)
}
