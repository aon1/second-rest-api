package repository

import "github.com/aon1/second-rest-api/model"

type BookRepository interface {
	FindAll() ([] model.Book, error)
}

type AuthorRepository interface {
	FindAll() ([] model.Author, error)
}
