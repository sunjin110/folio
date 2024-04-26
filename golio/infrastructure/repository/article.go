package repository

import (
	"context"
	"fmt"
	"log/slog"

	_ "embed"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
)

//go:embed query/create_article_bodies.sql
var createArticleBodiesSQL string

//go:embed query/create_article_summaries.sql
var createArticleSummariesSQL string

type article struct {
	d1Client d1.Client
}

func NewArticle(ctx context.Context, d1Client d1.Client) (repository.Article, error) {
	article := &article{
		d1Client: d1Client,
	}
	if err := article.creteTables(ctx); err != nil {
		return nil, fmt.Errorf("failed article.createTables: %w", err)
	}
	return article, nil
}

func (a *article) creteTables(ctx context.Context) error {
	slog.Info("create article tables...")
	_, err := a.d1Client.Query(ctx, &d1.Input{
		SQL: createArticleBodiesSQL,
	})
	if err != nil {
		return fmt.Errorf("failed create article_bodies. err: %w", err)
	}

	if _, err = a.d1Client.Query(ctx, &d1.Input{
		SQL: createArticleSummariesSQL,
	}); err != nil {
		return fmt.Errorf("failed create article summaries. err: %w", err)
	}
	return nil
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
