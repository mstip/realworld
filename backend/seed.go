package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Seed(db *sqlx.DB) {
	tags := []string{"tageins", "tagszwei", "tagdrei"}
	for _, t := range tags {
		db.MustExec(fmt.Sprintf(`INSERT INTO tags(tag) values("%s")`, t))
	}
}
