package main

import "testing"

func TestCreateArticle(t *testing.T) {
	db := NewSeededInMemDB()
	as := NewArticleService(db)
	auth := NewAuthService(db)
	u, err := auth.Register("user1", "user@user.de", "a1b2c3d4")
	if err != nil {
		t.Fatalf("register err %#v", err)
	}

	title := "test title moep"
	description := "this is an awesome test article"
	body := `this is an awesome test article!!! woooaahhh
	awesome
	!!!!
	moep
	`

	tagList := []string{"test", "article", "awesome"}

	a, err := as.CreateArticle(u.ID, title, description, body, tagList)
	if err != nil {
		t.Fatalf("create article %#v", err)
	}

	if a.AuthorID != u.ID {
		t.Errorf("authorID should %#v but is %#v", u.ID, a.AuthorID)
	}

	if a.Title != title {
		t.Errorf("Title should %#v but is %#v", title, a.Title)
	}

	if a.Description != description {
		t.Errorf("description should %#v but is %#v", description, a.Description)
	}

	if a.Body != body {
		t.Errorf("body should %#v but is %#v", body, a.AuthorID)
	}

	if a.Author.Username != u.Username {
		t.Errorf("Author.Username should %#v but is %#v", u.Username, a.Author.Username)
	}

	if len(a.TagList) != len(tagList) {
		t.Errorf("len(TagList)  should %#v but is %#v", len(tagList), len(a.TagList))
	}

	if a.Slug == "" || a.Slug == title {
		t.Errorf("Slug should not be empty or equal to tile %#v ", a.Slug)
	}

	if a.CreatedAt == "" {
		t.Errorf("CreatedAt should not be empty ")
	}
	if a.UpdatedAt == "" {
		t.Errorf("UpdatedAt should not be empty ")
	}
}

func TestAllArticles(t *testing.T) {
	db := NewSeededInMemDB()
	as := NewArticleService(db)
	articles, err := as.AllArticles()
	if err != nil {
		t.Fatalf("all article %#v", err)
	}

	if len(articles) == 0 {
		t.Errorf("len(articles) should not be 0")
	}
}
