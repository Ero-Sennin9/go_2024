package main

import (
	"fmt"
	"strconv"
	"task1/library"
)

type SliceStorage struct {
	books   []library.Book
	booksId []string
}

func (sliceStorage SliceStorage) Search(id string) (library.Book, bool) {
	for ind, currId := range sliceStorage.booksId {
		if currId == id {
			return sliceStorage.books[ind], true
		}
	}
	return library.Book{}, false
}

func (sliceStorage *SliceStorage) Add(id string, book library.Book) {
	sliceStorage.books = append(sliceStorage.books, book)
	sliceStorage.booksId = append(sliceStorage.booksId, id)
}

func (sliceStorage *SliceStorage) ChangeId(id library.Id) {
	newBooksId := []string{}
	for _, book := range sliceStorage.books {
		newBooksId = append(newBooksId, id(book))
	}
	sliceStorage.booksId = newBooksId
}

type MapStorage struct {
	data map[string]library.Book
}

func (mapStorage MapStorage) Search(id string) (library.Book, bool) {
	book, result := mapStorage.data[id]
	return book, result
}

func (mapStorage *MapStorage) Add(id string, book library.Book) {
	if mapStorage.data == nil {
		mapStorage.data = make(map[string]library.Book)
	}
	mapStorage.data[id] = book
}

func (mapStorage *MapStorage) ChangeId(id library.Id) {
	newData := make(map[string]library.Book)
	for _, book := range mapStorage.data {
		newData[id(book)] = book
	}
	mapStorage.data = newData
}

func Id1(book library.Book) string {
	return book.Name
}

func Id2(book library.Book) string {
	return book.Name + "_" + strconv.Itoa(len(book.Name))
}

func main() {
	books := []library.Book{library.Book{Name: "Book1"}, library.Book{Name: "Book2"}, library.Book{Name: "Book3"},
		library.Book{Name: "Book4"}, library.Book{Name: "Book5"}}

	libraryExample := library.Library{Storage: new(SliceStorage), Id: Id1}
	for _, book := range books {
		libraryExample.Add(book)
	}
	fmt.Println(libraryExample.Search("Book1"))
	fmt.Println(libraryExample.Search("Book3"))
	libraryExample.SetId(Id2)
	fmt.Println(libraryExample.Search("Book1"))
	fmt.Println(libraryExample.Search("Book3"))
	libraryExample.SetId(Id1)
	fmt.Println(libraryExample.Search("Book1"))
	fmt.Println(libraryExample.Search("Book3"))

	libraryExample.Storage = new(MapStorage)
	for _, book := range books {
		libraryExample.Add(book)
	}
	fmt.Println(libraryExample.Search("Book1"))
	fmt.Println(libraryExample.Search("Book3"))
	libraryExample.SetId(Id2)
	fmt.Println(libraryExample.Search("Book1"))
	fmt.Println(libraryExample.Search("Book3"))
	libraryExample.SetId(Id1)
	fmt.Println(libraryExample.Search("Book1"))
	fmt.Println(libraryExample.Search("Book3"))
}
