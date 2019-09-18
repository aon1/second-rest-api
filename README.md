Simple rest api written in go

Create a database called bookstore on your postgres server

Restore database
```
psql dbname < dump.sql
```
Install
```
go install
```
Run
```
go run main.go
```

Database config can be changed on main.go file

http://localhost:3000/authors or http://localhost:3000/books
