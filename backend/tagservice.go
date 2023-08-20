package main

import (
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
