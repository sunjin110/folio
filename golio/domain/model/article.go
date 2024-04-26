package model

import "time"

// Article 記事
type Article struct {
	ID        string
	Title     string
	Body      string
	Writer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ArticleSummary 記事の概要
type ArticleSummary struct {
	ID        string
	Title     string
	Writer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
