package usecase

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Article interface {
	Get(ctx context.Context, id string) (*model.Article, error)
	Insert(ctx context.Context, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context)
}

type article struct {
}

func NewArticle() Article {
	return &article{}
}

func (a *article) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (a *article) Find(ctx context.Context) {
	panic("unimplemented")
}

func (a *article) Get(ctx context.Context, id string) (*model.Article, error) {
	panic("unimplemented")
}

func (a *article) Insert(ctx context.Context, article *model.Article) error {
	// 認証/認可

	panic("unimplemented")
}

func (a *article) Update(ctx context.Context, article *model.Article) error {
	panic("unimplemented")
}
