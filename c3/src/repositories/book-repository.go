package repositories

import (
	"c3/src/models"
	"fmt"
)

type BookRepo struct {
	bookList []models.Book
}

func (bookRepo *BookRepo) AddBook(book models.Book) error {
	bookRepo.bookList = append(bookRepo.bookList, book)
	return nil
}

func (bookRepo *BookRepo) GetBooks() ([]models.Book, error) {
	return bookRepo.bookList, nil
}

func (bookRepo *BookRepo) GetBook(title string) (models.Book, error) {
	for _, book := range bookRepo.bookList {
		if book.Title == title {
			return book, nil
		}
	}
	// Aqui se retorna el caso en el que no se encontro por lo cual hay un error
	// Esto es un libro "vacio", ya que no se puede devolver null
	// junto con el error
	return models.Book{}, fmt.Errorf("Book not found: %s", title)
}
