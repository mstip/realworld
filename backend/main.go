package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSeededInMemDB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	Migrate(db)
	Seed(db)
	return db
}

func main() {
	db := NewSeededInMemDB()
	r := NewRouter(db)
	http.Handle("/", r)

	log.Println("up and running!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
