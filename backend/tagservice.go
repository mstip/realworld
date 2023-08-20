package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Tag string

type DBTag struct {
	ID  int    `db:"id"`
	Tag string `db:"tag"`
}

type TagService struct {
	db *sqlx.DB
}

func NewTagService(db *sqlx.DB) *TagService {
	return &TagService{
		db: db,
	}
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

func (ts *TagService) CreateTag(tag string) (DBTag, error) {
	_, err := ts.db.Exec("INSERT OR IGNORE INTO tags(tag) values(?)", tag)
	if err != nil {
		return DBTag{}, err
	}

	var dbTags []DBTag

	if err := ts.db.Select(&dbTags, "SELECT * FROM tags WHERE tag = ?", tag); err != nil {
		return DBTag{}, err
	}

	if len(dbTags) == 0 {
		return DBTag{}, fmt.Errorf("cant find tag %s", tag)
	}

	return dbTags[0], nil
}
