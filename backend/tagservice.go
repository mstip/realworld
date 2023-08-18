package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Tag string

type TagService struct {
	db *sqlx.DB
}

func NewTagService(db *sqlx.DB) (*TagService, error) {
	return &TagService{
		db: db,
	}, nil
}

func (ts *TagService) Migrate() error {
	ts.db.MustExec(`
	CREATE TABLE IF NOT EXISTS"tags" (
		"id"	INTEGER NOT NULL,
		"tag"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)
	return nil
}

func (ts *TagService) Seed() error {
	tags := []string{"tageins", "tagszwei", "tagdrei"}
	for _, t := range tags {
		ts.db.MustExec(fmt.Sprintf(`INSERT INTO tags(tag) values("%s")`, t))
	}
	return nil
}

func (ts *TagService) AllTags() ([]string, error) {
	var tags []string

	if err := ts.db.Select(&tags, "SELECT tag FROM tags"); err != nil {
		return nil, err
	}

	if tags == nil {
		tags = []string{}
	}

	return tags, nil
}
