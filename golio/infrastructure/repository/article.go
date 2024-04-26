package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
)

type article struct {
	d1Client d1.Client
}

func NewArticle(d1Client d1.Client) (repository.Article, error) {
	return &article{
		d1Client: d1Client,
	}, nil
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
