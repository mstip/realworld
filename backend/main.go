package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlx.MustConnect("sqlite3", ":memory:")

	tagService, _ := NewTagService(db)
	tagService.Migrate()
	tagService.Seed()

	r := NewRouter(db)
	http.Handle("/", r)

	log.Println("up and running!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
