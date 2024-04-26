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

//go:embed query/find_one_article_summaries.sql
var findOneArticleSummariesSQL string

//go:embed query/find_one_article_bodies.sql
var findOneArticleBodiesSQL string

//go:embed query/upsert_article_bodies.sql
var upsertArticleBodiesSQL string

//go:embed query/upsert_article_summaries.sql
var upsertArticleSummariesSQL string

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
	result, err := a.d1Client.Query(ctx, &d1.Input{
		Params: []string{id},
		SQL:    findOneArticleSummariesSQL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed d1Client.Query. err: %w", err)
	}
	// result, err := a.d1Client.Query(ctx, &d1.Input{
	// 	Params: []string{id},
	// 	SQL:    findOneArticleBodiesSQL,
	// })

	fmt.Println("result is ", result)
	return nil, nil
}

func (a *article) Insert(ctx context.Context, article *model.Article) error {
	_, err := a.d1Client.Query(ctx, &d1.Input{
		Params: []string{article.ID, article.ID, article.Body, fmt.Sprintf("%d", article.CreatedAt.Unix()),
			fmt.Sprintf("%d", article.UpdatedAt.Unix())},
		SQL: upsertArticleBodiesSQL,
	})
	if err != nil {
		return fmt.Errorf("failed insert articleBody. article: %+v, err: %w", article, err)
	}

	_, err = a.d1Client.Query(ctx, &d1.Input{
		Params: []string{article.ID, article.Title, fmt.Sprintf("%d", article.CreatedAt.Unix()), fmt.Sprintf("%d", article.UpdatedAt.Unix())},
		SQL:    upsertArticleSummariesSQL,
	})
	if err != nil {
		return fmt.Errorf("failed insert articleSummary. article: %+v, err: %w", article, err)
	}
	return nil
}

func (a *article) Update(ctx context.Context, article *model.Article) error {
	return a.Insert(ctx, article)
}

func (a *article) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}
