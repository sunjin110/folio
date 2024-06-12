package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type ArticleAI interface {
	ChangeBodyByAI(ctx context.Context, article *model.Article, orderToAI string) (*model.Article, error)
	GenerateBodyByAI(ctx context.Context, prompt string) (string, error)
}
