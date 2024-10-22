package book

type Book struct {
	Name string
}

type Id func(Book) string
