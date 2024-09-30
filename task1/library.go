package main

type Library struct {
	storage     Storage
	id_function func(string) int
}

func (library *Library) addBook(book Book) {
	var book_id = library.id_function(book.name)
	library.storage.addBook(book_id, book)
}

func (library *Library) getBook(name string) (Book, bool) {
	var book_id = library.id_function(name)
	return library.storage.getBook(book_id)
}

func (library *Library) changeIdFunction(id_function func(string) int) {
	library.id_function = id_function
}

func (library *Library) changeStorage(storage Storage) {
	library.storage = storage
}
