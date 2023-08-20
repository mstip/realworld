package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
)

// TODO: favorite
// TODO: follower

type Article struct {
	ID             int      `json:"id"`
	Slug           string   `json:"slug"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Body           string   `json:"body"`
	TagList        []string `json:"tagList"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	AuthorID       int      `json:"authorID"`
	Favorited      bool     `json:"favorited"`
	FavoritesCount int      `json:"favoritesCount"`
	Author         struct {
		Username  string `json:"username"`
		Bio       string `json:"nio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"author"`
}

type DBArticle struct {
	ID          int
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	AuthorID    int    `db:"author_id"`
}

type DBArticleWithUser struct {
	ID          int
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	AuthorID    int    `db:"author_id"`
	Username    string
	Bio         sql.NullString
	Image       sql.NullString
}

type ArticleService struct {
	db *sqlx.DB
}

func NewArticleService(db *sqlx.DB) *ArticleService {
	return &ArticleService{
		db: db,
	}
}

func (as *ArticleService) AllArticles() ([]Article, error) {
	//
	var dbArticles []DBArticleWithUser
	err := as.db.Select(&dbArticles, `
	SELECT a.id, a.slug, a.description, a.body, a.title, a.created_at, a.updated_at, a.author_id, u.username, u.bio, u.image 
	FROM articles as a 
	INNER JOIN users as u ON u.id = a.author_id;`,
	)
	if err != nil {
		return []Article{}, err
	}

	articles := []Article{}

	for _, dbArticle := range dbArticles {

		var tagList []string

		err = as.db.Select(&tagList,
			`SELECT t.tag FROM tags as t INNER JOIN articles_tags as at ON at.tag_id = t.id WHERE at.article_id = ?`,
			dbArticle.ID,
		)

		if err != nil {
			return []Article{}, err
		}

		articles = append(articles, Article{
			ID:             dbArticle.ID,
			Slug:           dbArticle.Slug,
			Title:          dbArticle.Title,
			Description:    dbArticle.Description,
			Body:           dbArticle.Body,
			TagList:        tagList,
			CreatedAt:      dbArticle.CreatedAt,
			UpdatedAt:      dbArticle.UpdatedAt,
			AuthorID:       dbArticle.AuthorID,
			Favorited:      false,
			FavoritesCount: 0,
			Author: struct {
				Username  string `json:"username"`
				Bio       string `json:"nio"`
				Image     string `json:"image"`
				Following bool   `json:"following"`
			}{
				Username:  dbArticle.Username,
				Bio:       dbArticle.Bio.String,
				Image:     dbArticle.Image.String,
				Following: false,
			},
		})
	}

	return articles, nil
}

func (as *ArticleService) CreateArticle(authorID int, title string, description string, body string, tagList []string) (Article, error) {
	res, err := as.db.Exec(
		`INSERT INTO articles(slug, title, description, body, author_id, created_at, updated_at) values(?,?,?,?,?,?,?)`,
		slug.Make(title), title, description, body, authorID, time.Now(), time.Now(),
	)

	if err != nil {
		return Article{}, err
	}

	articleID, _ := res.LastInsertId()

	var dbArticles []DBArticle
	err = as.db.Select(&dbArticles, "SELECT * FROM articles WHERE id = ?", articleID)
	if err != nil {
		return Article{}, err
	}

	if len(dbArticles) == 0 {
		return Article{}, fmt.Errorf("cant find new article with id %d", articleID)
	}

	var u []User

	err = as.db.Select(&u, "SELECT * FROM users WHERE id = ?", authorID)
	if err != nil {
		return Article{}, err
	}

	if len(u) == 0 {
		return Article{}, fmt.Errorf("failed to get author with id %d", authorID)
	}

	ts := NewTagService(as.db)

	for _, t := range tagList {
		dbTag, err := ts.CreateTag(t)
		if err != nil {
			return Article{}, err
		}

		_, err = as.db.Exec(`INSERT INTO articles_tags(article_id, tag_id) values(?,?)`, dbArticles[0].ID, dbTag.ID)
		if err != nil {
			return Article{}, err
		}
	}

	a := Article{
		ID:             dbArticles[0].ID,
		Slug:           dbArticles[0].Slug,
		Title:          dbArticles[0].Title,
		Description:    dbArticles[0].Description,
		Body:           dbArticles[0].Body,
		TagList:        tagList,
		CreatedAt:      dbArticles[0].CreatedAt,
		UpdatedAt:      dbArticles[0].UpdatedAt,
		AuthorID:       dbArticles[0].AuthorID,
		Favorited:      false,
		FavoritesCount: 0,
		Author: struct {
			Username  string `json:"username"`
			Bio       string `json:"nio"`
			Image     string `json:"image"`
			Following bool   `json:"following"`
		}{
			Username:  u[0].Username,
			Bio:       u[0].Bio.String,
			Image:     u[0].Image.String,
			Following: false,
		},
	}

	return a, nil
}
