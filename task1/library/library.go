package library

type Id func(Book) string

type Book struct {
	Name string
}

type Storage interface {
	Search(id string) (Book, bool)
	Add(id string, book Book)
	ChangeId(id Id)
}

type Library struct {
	Storage
	Id Id
}

func (lib Library) Search(name string) (Book, bool) {
	return lib.Storage.Search(lib.Id(Book{Name: name}))
}

func (lib *Library) Add(book Book) {
	lib.Storage.Add(lib.Id(book), book)
}

func (lib *Library) SetId(id Id) {
	lib.Storage.ChangeId(id)
	lib.Id = id
}
