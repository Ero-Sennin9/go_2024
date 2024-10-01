package slice_storage

import "task1/library/book"

type SliceStorage struct {
	books   []book.Book
	booksId []string
}

func (sliceStorage SliceStorage) Search(id string) (book.Book, bool) {
	for ind, currId := range sliceStorage.booksId {
		if currId == id {
			return sliceStorage.books[ind], true
		}
	}
	return book.Book{}, false
}

func (sliceStorage *SliceStorage) Add(id string, book book.Book) {
	sliceStorage.books = append(sliceStorage.books, book)
	sliceStorage.booksId = append(sliceStorage.booksId, id)
}

func (sliceStorage *SliceStorage) ChangeId(id book.Id) {
	newBooksId := []string{}
	for _, book := range sliceStorage.books {
		newBooksId = append(newBooksId, id(book))
	}
	sliceStorage.booksId = newBooksId
}

func MakeSliceStorage() *SliceStorage {
	return new(SliceStorage)
}
