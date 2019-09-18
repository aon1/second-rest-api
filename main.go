package main

import (
	"fmt"
	"github.com/aon1/second-rest-api/database"
	"net/http"
	"os"

	"github.com/aon1/second-rest-api/handler"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "bookstore"
)

func main() {
	connection, err := database.ConnectPostgres(host, user, password, dbname, port)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	bookHandler := handler.NewBookHandler(connection)
	authorHandler := handler.NewAuthorHandler(connection)

	http.HandleFunc("/books", bookHandler.FindAll)
	http.HandleFunc("/authors", authorHandler.FindAll)
	http.ListenAndServe(":3000", nil)
}
