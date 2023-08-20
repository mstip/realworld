package main

import "github.com/jmoiron/sqlx"

func Migrate(db *sqlx.DB) {
	db.MustExec(`
	CREATE TABLE "users" (
		"id"	INTEGER NOT NULL,
		"email"	TEXT NOT NULL UNIQUE,
		"username"	TEXT NOT NULL UNIQUE,
		"bio"	TEXT,
		"image"	TEXT,
		"token"	TEXT,
		"password"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)
	`)
	db.MustExec(`
	CREATE TABLE "articles" (
		"id"	INTEGER NOT NULL,
		"slug"	TEXT NOT NULL UNIQUE,
		"description"	TEXT NOT NULL,
		"body"	TEXT NOT NULL,
		"title"	TEXT NOT NULL,
		"created_at"	TEXT NOT NULL,
		"updated_at"	TEXT NOT NULL,
		"author_id"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);
	`)
	db.MustExec(`
	CREATE TABLE "tags" (
		"id"	INTEGER NOT NULL,
		"tag"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)

	db.MustExec(`
	CREATE TABLE "articles_tags" (
		"id"	INTEGER NOT NULL,
		"article_id"	INTEGER NOT NULL,
		"tag_id"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)
	`)
}
