package repointerfaces

import "c7/src/models"

type BookRepo interface {
	AddBook(book models.Book) error
	GetBooks() ([]models.Book, error)
	GetBook(title string) (models.Book, error)
}
