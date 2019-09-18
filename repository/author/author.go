package author

import (
	"database/sql"
	"github.com/aon1/second-rest-api/model"
	"github.com/aon1/second-rest-api/repository"
)

type authorRepository struct {
	Conn *sql.DB
}

func NewAuthorRepo(Conn *sql.DB) repository.AuthorRepository {
	return &authorRepository{Conn:Conn}
}

func (t *authorRepository) FindAll() ([] model.Author, error) {
	rows, err := t.Conn.Query("SELECT * FROM author")
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var authors []model.Author

	for rows.Next() {
		author := model.Author{}
		err = rows.Scan(&author.ID, &author.Name)

		authors = append(authors, author)
	}

	return authors, err
}