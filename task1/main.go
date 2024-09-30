package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
	var books_first_slice []Book = []Book{
		{name: "Мастер и Маргарита", author: "Михаил Булгаков"},
		{name: "Отцы и дети", author: "Иван Тургенев"},
		{name: "Горе от ума", author: "Александр Грибоедов"},
	}
	var books_second_slice []Book = []Book{
		{name: "1984", author: "George Orwell"},
		{name: "The Little Prince", author: "Antoine de Saint-Exupery"},
	}
	var library Library = createLibrary(first_function)
	for _, book := range books_first_slice {
		library.addBook(book)
	}

	var val, ok = library.getBook("Отцы и дети")
	fmt.Println(val.author, ok) // Иван Тургенев true

	library.changeIdFunction(second_function)

	val, ok = library.getBook("Мастер и Маргарита")
	fmt.Println(val.author, ok) // Михаил Булгаков true

	library.clearStorage()
	for _, book := range books_second_slice {
		library.addBook(book)
	}

	val, ok = library.getBook("1984")
	fmt.Println(val.author, ok) // George Orwell true
}
