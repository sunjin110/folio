package model

import (
	"time"

	"github.com/google/uuid"
)

type ArticleTag struct {
	ID          string
	Name        string
	CreatedTime time.Time
	UpdatedTime time.Time
}

func NewArticleTag(name string, createdTime time.Time) *ArticleTag {
	uu, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return &ArticleTag{
		ID:          uu.String(),
		Name:        name,
		CreatedTime: createdTime,
		UpdatedTime: createdTime,
	}
}
