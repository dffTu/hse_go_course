package main

type Library struct {
	storage     Storage
	books_id    map[string]int
	id_function func(string) int
}

func (library *Library) addBook(book Book) {
	var book_id = library.id_function(book.name)
	library.books_id[book.name] = book_id
	library.storage.addBook(book_id, book)
}

func (library *Library) getBook(name string) (Book, bool) {
	var val, ok = library.books_id[name]
	if ok {
		return library.storage.getBook(val)
	}
	return Book{}, false
}

func (library *Library) changeIdFunction(id_function func(string) int) {
	library.id_function = id_function
}

func (library *Library) changeStorage(storage Storage) {
	library.storage = storage
}
