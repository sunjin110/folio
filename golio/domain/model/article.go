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
	Tags      []*ArticleTag
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (article *Article) GetTagIDs() []string {
	tagIDs := make([]string, 0, len(article.Tags))
	for _, tag := range article.Tags {
		tagIDs = append(tagIDs, tag.ID)
	}
	return tagIDs
}

func NewArticle(title string, body string, writer string, tags []*ArticleTag, createdAt time.Time) *Article {
	uu, err := uuid.NewRandom()
	if err != nil {
		// ほぼ発生しない
		panic(err)
	}
	return &Article{
		ID:        uu.String(),
		Title:     title,
		Body:      body,
		Tags:      tags,
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
	Tags      []*ArticleTag
	CreatedAt time.Time
	UpdatedAt time.Time
}
