package handler

import (
	"github.com/aon1/second-rest-api/database"
	"github.com/aon1/second-rest-api/repository"
	"github.com/aon1/second-rest-api/repository/author"
	"github.com/aon1/second-rest-api/helper"
	"net/http"
)

func NewAuthorHandler(db *database.DB) *Author {
	return &Author{repo: author.NewAuthorRepo(db.SQL)}
}

type Author struct {
	repo repository.AuthorRepository
}

func (t *Author) FindAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := t.repo.FindAll()

	helper.RespondWithJSON(w, http.StatusOK, payload)
}

