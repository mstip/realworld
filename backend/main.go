package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	var counter []int

	err := db.Select(&counter, "SELECT number FROM counter limit 1")
	if err != nil || len(counter) == 0 {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	counter[0]++
	db.MustExec(fmt.Sprintf("UPDATE counter SET number = %d", counter[0]))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HOME! count %d", counter[0])
}

func main() {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec("CREATE TABLE counter(number int);")
	db.MustExec("INSERT INTO counter(number) values(0)")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, db)
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", r))
}
