package main

import (
	"github.com/jmoiron/sqlx"
)

func Seed(db *sqlx.DB) {
	tagService := NewTagService(db)
	authService := NewAuthService(db)
	articleService := NewArticleService(db)

	tags := []string{"tageins", "tagszwei", "tagdrei"}
	for _, t := range tags {
		tagService.CreateTag(t)
	}

	authService.Register("testuser1", "testuser1@user.de", "a1b2c3d4")
	authService.Register("testuser2", "testuser2@user.de", "a1b2c3d4")
	authService.Register("testuser3", "testuser3@user.de", "a1b2c3d4")

	articleService.CreateArticle(1, "test article", "test description", "lorem ipsum", []string{"tageins"})
	articleService.CreateArticle(1, "test article2", "test description2", "lorem ipsum2", []string{"tagzwei"})
	articleService.CreateArticle(1, "test article3", "test description3", "lorem ipsum3", []string{"tagdrei"})
}
