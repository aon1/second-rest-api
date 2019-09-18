package book

import (
	"database/sql"
	"github.com/aon1/second-rest-api/model"
	"github.com/aon1/second-rest-api/repository"
)

type bookRepository struct {
	Conn *sql.DB
}

func NewBookRepo(Conn *sql.DB) repository.BookRepository {
	return &bookRepository{Conn:Conn}
}

func (t *bookRepository) FindAll() ([] model.Book, error) {
	rows, err := t.Conn.Query("SELECT b.isbn, b.title, a.name as author, b.price FROM books b JOIN author a ON a.id = b.author_id")
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var books []model.Book

	for rows.Next() {
		book := model.Book{}
		err = rows.Scan(&book.ISBN, &book.Title, &book.Author, &book.Price)

		books = append(books, book)
	}

	return books, err
}