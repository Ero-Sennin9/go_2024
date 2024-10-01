package main

import (
	"fmt"
	"strconv"
	"task1/library"
	"task1/library/book"
	"task1/library/storage/map_storage"
	"task1/library/storage/slice_storage"
)

func Id1(book book.Book) string {
	return book.Name
}

func Id2(book book.Book) string {
	return book.Name + "_" + strconv.Itoa(len(book.Name))
}

func main() {
	books := []book.Book{book.Book{Name: "Book1"}, book.Book{Name: "Book2"}, book.Book{Name: "Book3"},
		book.Book{Name: "Book4"}, book.Book{Name: "Book5"}}

	libraryExample := library.Library{Storage: new(slice_storage.SliceStorage), Id: Id1}
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

	libraryExample.Storage = new(map_storage.MapStorage)
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
