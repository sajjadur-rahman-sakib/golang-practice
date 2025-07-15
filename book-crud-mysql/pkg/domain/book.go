package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

type IBookRepo interface {
	GetBooks(bookID uint) []models.Book
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}

type IBookService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}
