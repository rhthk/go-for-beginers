package module4

type Book struct {
	name   string
	author string
	cost   float64
}

type Library struct {
	books []Book
}

func (lib *Library) addBook(book Book) {
	lib.books = append(lib.books, book)
}

func (lib *Library) totalValue() float64 {
	total := 0.0
	for _, book := range lib.books {
		total = total + book.cost
	}
	return total
}
