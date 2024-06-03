package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	_ "embed"

	"github.com/huandu/go-sqlbuilder"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
	"github.com/sunjin110/folio/golio/infrastructure/repository/conv"
	"golang.org/x/sync/errgroup"
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
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("article_summaries").Limit(paging.Limit).Offset(paging.Offset).OrderBy("created_at")
	if sortType == repository.SortTypeAsc {
		sb.Asc()
	} else {
		sb.Desc()
	}
	sql, args := sb.Build()
	output, err := a.d1Client.Query(ctx, &d1.Input{
		Params: args,
		SQL:    sql,
	})
	if err != nil {
		return nil, fmt.Errorf("failed find summary. sql: %s, err: %w", sql, err)
	}
	return conv.ToArticleSummaries(output), nil
}

func (a *article) Get(ctx context.Context, id string) (*model.Article, error) {

	eg, ctx := errgroup.WithContext(ctx)

	var summaries []map[string]interface{}
	eg.Go(func() error {
		summariesOutput, err := a.d1Client.Query(ctx, &d1.Input{
			Params: []interface{}{id},
			SQL:    findOneArticleSummariesSQL,
		})
		if err != nil {
			return fmt.Errorf("failed summaries d1Client.Query. err: %w", err)
		}
		if len(summariesOutput.Results) == 0 {
			return repository.ErrNotFound
		}

		summaries = summariesOutput.GetResultMapList()
		return nil
	})

	var bodies []map[string]interface{}
	eg.Go(func() error {
		bodiesOutput, err := a.d1Client.Query(ctx, &d1.Input{
			Params: []interface{}{id},
			SQL:    findOneArticleBodiesSQL,
		})
		if err != nil {
			return fmt.Errorf("failed bodies d1Client.Query. err: %w", err)
		}
		if len(bodiesOutput.Results) == 0 {
			return repository.ErrNotFound
		}

		bodies = bodiesOutput.GetResultMapList()
		return nil
	})

	if err := eg.Wait(); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed article.Get: %w", err)
	}

	return &model.Article{
		ID:        id,
		Title:     summaries[0]["title"].(string),
		Body:      bodies[0]["body"].(string),
		Writer:    "",
		CreatedAt: time.Unix(int64(summaries[0]["created_at"].(float64)), 0),
		UpdatedAt: time.Unix(int64(summaries[0]["updated_at"].(float64)), 0),
	}, nil
}

func (a *article) Insert(ctx context.Context, article *model.Article) error {
	_, err := a.d1Client.Query(ctx, &d1.Input{
		Params: []interface{}{article.ID, article.ID, article.Body, fmt.Sprintf("%d", article.CreatedAt.Unix()),
			fmt.Sprintf("%d", article.UpdatedAt.Unix())},
		SQL: upsertArticleBodiesSQL,
	})
	if err != nil {
		return fmt.Errorf("failed insert articleBody. article: %+v, err: %w", article, err)
	}

	_, err = a.d1Client.Query(ctx, &d1.Input{
		Params: []interface{}{article.ID, article.Title, fmt.Sprintf("%d", article.CreatedAt.Unix()), fmt.Sprintf("%d", article.UpdatedAt.Unix())},
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
	if _, err := a.d1Client.Query(ctx, &d1.Input{
		Params: []interface{}{id},
		SQL:    `delete from article_summaries where id = ?`,
	}); err != nil {
		return fmt.Errorf("failed delete article_summaries. id: %s, err: %w", id, err)
	}

	if _, err := a.d1Client.Query(ctx, &d1.Input{
		Params: []interface{}{id},
		SQL:    `delete from article_bodies where article_summaries_id = ?`,
	}); err != nil {
		return fmt.Errorf("failed delete article_bodies. id: %s, err: %w", id, err)
	}
	return nil
}

func (a *article) CountTotal(ctx context.Context) (int32, error) {
	sb := sqlbuilder.NewSelectBuilder().Select("count(*)").From("article_summaries")

	sql, args := sb.Build()
	output, err := a.d1Client.Query(ctx, &d1.Input{
		Params: args,
		SQL:    sql,
	})
	if err != nil {
		return -1, fmt.Errorf("failed count article_summaries. sql: %s, err: %w", sql, err)
	}
	countTotal := int32(output.GetResultMapList()[0]["count(*)"].(float64))
	return countTotal, nil
}
