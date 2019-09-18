package handler

import (
	"github.com/aon1/second-rest-api/database"
	"github.com/aon1/second-rest-api/helper"
	"github.com/aon1/second-rest-api/repository"
	"github.com/aon1/second-rest-api/repository/book"
	"net/http"
)

func NewBookHandler(db *database.DB) *Book {
	return &Book{repo: book.NewBookRepo(db.SQL)}
}

type Book struct {
	repo repository.BookRepository
}

func (t *Book) FindAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := t.repo.FindAll()

	helper.RespondWithJSON(w, http.StatusOK, payload)
}
