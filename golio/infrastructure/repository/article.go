package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type article struct {
}

func NewArticle() repository.Article {
	return &article{}
}

func (a *article) FindSummary(ctx context.Context, sortType repository.SortType, paging *repository.Paging) ([]*model.ArticleSummary, error) {
	panic("unimplemented")
}

func (a *article) Get(ctx context.Context, id string) (*model.Article, error) {
	panic("unimplemented")
}

func (a *article) Insert(ctx context.Context, article *model.Article) error {
	panic("unimplemented")
}

func (a *article) Update(ctx context.Context, article *model.Article) error {
	panic("unimplemented")
}

func (a *article) Delete(ctx context.Context, id string) {
	panic("unimplemented")
}
