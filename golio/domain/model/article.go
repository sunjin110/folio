package model

import (
	"time"

	"github.com/google/uuid"
)

// Article 記事
type Article struct {
	ID        string
	Title     string
	Body      string
	Writer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(title string, body string, writer string, createdAt time.Time) *Article {
	uu, err := uuid.NewRandom()
	if err != nil {
		// ほぼ発生しない
		panic(err)
	}
	return &Article{
		ID:        uu.String(),
		Title:     title,
		Body:      body,
		Writer:    writer,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
}

// ArticleSummary 記事の概要
type ArticleSummary struct {
	ID        string
	Title     string
	Writer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
