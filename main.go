package main

import (
	"fmt"
	"strconv"
)

type Id func(Book) string

type Book struct {
	name string
}

type Storage interface {
	Search(id string) (Book, bool)
	Add(id string, book Book)
	ChangeId(id Id)
}

type SliceStorage struct {
	books   []Book
	booksId []string
}

func (sliceStorage SliceStorage) Search(id string) (Book, bool) {
	for ind, currId := range sliceStorage.booksId {
		if currId == id {
			return sliceStorage.books[ind], true
		}
	}
	return Book{}, false
}

func (sliceStorage *SliceStorage) Add(id string, book Book) {
	sliceStorage.books = append(sliceStorage.books, book)
	sliceStorage.booksId = append(sliceStorage.booksId, id)
}

func (sliceStorage *SliceStorage) ChangeId(id Id) {
	newBooksId := []string{}
	for _, book := range sliceStorage.books {
		newBooksId = append(newBooksId, id(book))
	}
	sliceStorage.booksId = newBooksId
}

type MapStorage struct {
	data map[string]Book
}

func (mapStorage MapStorage) Search(id string) (Book, bool) {
	book, result := mapStorage.data[id]
	return book, result
}

func (mapStorage *MapStorage) Add(id string, book Book) {
	if mapStorage.data == nil {
		mapStorage.data = make(map[string]Book)
	}
	mapStorage.data[id] = book
}

func (mapStorage *MapStorage) ChangeId(id Id) {
	newData := make(map[string]Book)
	for _, book := range mapStorage.data {
		newData[id(book)] = book
	}
	mapStorage.data = newData
}

type Library struct {
	Storage
	id Id
}

func (lib Library) Search(name string) (Book, bool) {
	return lib.Storage.Search(lib.id(Book{name: name}))
}

func (lib *Library) Add(book Book) {
	lib.Storage.Add(lib.id(book), book)
}

func (lib *Library) SetId(id Id) {
	lib.Storage.ChangeId(id)
	lib.id = id
}

func Id1(book Book) string {
	return book.name
}

func Id2(book Book) string {
	return book.name + "_" + strconv.Itoa(len(book.name))
}

func main() {
	books := []Book{Book{name: "Book1"}, Book{name: "Book2"}, Book{name: "Book3"},
		Book{name: "Book4"}, Book{name: "Book5"}}

	library := Library{Storage: new(SliceStorage), id: Id1}
	for _, book := range books {
		library.Add(book)
	}
	fmt.Println(library.Search("Book1"))
	fmt.Println(library.Search("Book3"))
	library.SetId(Id2)
	fmt.Println(library.Search("Book1"))
	fmt.Println(library.Search("Book3"))
	library.SetId(Id1)
	fmt.Println(library.Search("Book1"))
	fmt.Println(library.Search("Book3"))

	library.Storage = new(MapStorage)
	for _, book := range books {
		library.Add(book)
	}
	fmt.Println(library.Search("Book1"))
	fmt.Println(library.Search("Book3"))
	library.SetId(Id2)
	fmt.Println(library.Search("Book1"))
	fmt.Println(library.Search("Book3"))
	library.SetId(Id1)
	fmt.Println(library.Search("Book1"))
	fmt.Println(library.Search("Book3"))
}
