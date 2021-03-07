package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	conexao := "user=alandel dbname=alan_loja password=alandel host=localhost port=5051 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
