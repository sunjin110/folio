package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Article interface {
	Get(ctx context.Context, id string) (*model.Article, error)
	Insert(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, article *model.Article) error
	FindSummary(ctx context.Context, sortType SortType, paging *Paging, search *ArticleSearch) ([]*model.ArticleSummary, error)
	CountTotal(ctx context.Context, search *ArticleSearch) (int32, error)
	ChangeBodyByAI(ctx context.Context, article *model.Article, orderToAI string) (*model.Article, error)
}

type ArticleSearch struct {
	Title *string
}
