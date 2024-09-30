package main

type Storage struct {
	books map[int]Book
}

func (storage *Storage) addBook(id int, book Book) {
	(*storage).books[id] = book
}

func (storage *Storage) getBook(id int) (Book, bool) {
	book, ok := (*storage).books[id]
	if ok {
		delete((*storage).books, id)
	}
	return book, ok
}

func createStorage() Storage {
	return Storage{make(map[int]Book)}
}
