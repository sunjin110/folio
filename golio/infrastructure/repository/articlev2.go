package repository

import (
	"context"
	"fmt"
	"log/slog"

	_ "embed"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
)

var (
	//go:embed query/postgresql/upsert_article_summary.sql
	upsertArticleSummaryPostgreSQL string

	//go:embed query/postgresql/upsert_article_body.sql
	upsertArticleBodyPostgreSQL string
)

type articleV2 struct {
	db *sqlx.DB
}

func NewArticleV2(ctx context.Context, db *sqlx.DB) repository.Article {
	return &articleV2{
		db: db,
	}
}

func (a *articleV2) CountTotal(ctx context.Context) (int32, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)").From("article_summaries")

	sql, _ := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := a.db.QueryxContext(ctx, sql)
	if err != nil {
		return -1, fmt.Errorf("failed QueryContext. sql: %s, err: %w", sql, err)
	}

	if !rows.Next() {
		return -1, fmt.Errorf("not found rows. sql: %s", sql)
	}

	var totalCount int32
	if err := rows.Scan(&totalCount); err != nil {
		return -1, fmt.Errorf("failed scan. sql: %s, err: %w", sql, err)
	}
	return totalCount, nil
}

func (a *articleV2) Delete(ctx context.Context, id string) (err error) {
	tx, err := a.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed begin. err: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				slog.ErrorContext(ctx, "failed article delete rollback", "err", err, "id", id)
			}
		}
	}()

	if _, err := tx.ExecContext(ctx, `delete from article_bodies where id = $1`, id); err != nil {
		return fmt.Errorf("failed delete from article_bodies. id: %s, err: %w", id, err)
	}

	if _, err := tx.ExecContext(ctx, "delete from article_summaries where id = $1", id); err != nil {
		return fmt.Errorf("failed delete from article_summaries. id: %s, err: %w", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed delete commit. err: %w", err)
	}
	return nil
}

func (a *articleV2) FindSummary(ctx context.Context, sortType repository.SortType, paging *repository.Paging) ([]*model.ArticleSummary, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("article_summaries").Limit(paging.Limit).Offset(paging.Offset).OrderBy("created_at")
	if sortType == repository.SortTypeAsc {
		sb.Asc()
	} else {
		sb.Desc()
	}
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	summaries := postgres_dto.ArticleSummaries{}
	if err := a.db.SelectContext(ctx, &summaries, sql, args...); err != nil {
		return nil, fmt.Errorf("failed get select article_summaries. sql: %s, err: %w", sql, err)
	}

	return summaries.ToSummariesModel(), nil
}

func (a *articleV2) Get(ctx context.Context, id string) (*model.Article, error) {
	body := &postgres_dto.ArticleBody{}
	if err := a.db.GetContext(ctx, body, "select * from article_bodies where id = $1", id); err != nil {
		return nil, fmt.Errorf("failed get article_bodies. err: %w", err)
	}

	summary := &postgres_dto.ArticleSummary{}
	if err := a.db.GetContext(ctx, summary, "select * from article_summaries where id = $1", id); err != nil {
		return nil, fmt.Errorf("failed get article_summaries. err: %w", err)
	}

	return &model.Article{
		ID:        summary.ID,
		Title:     summary.Title,
		Body:      body.Body,
		Writer:    "",
		CreatedAt: summary.CreatedAt,
		UpdatedAt: summary.UpdatedAt,
	}, nil
}

func (a *articleV2) Insert(ctx context.Context, article *model.Article) error {
	return a.upsert(ctx, article)
}

func (a *articleV2) Update(ctx context.Context, article *model.Article) error {
	return a.upsert(ctx, article)
}

func (a *articleV2) upsert(ctx context.Context, article *model.Article) (err error) {

	tx, err := a.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed upsert transaction begin. err: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				slog.ErrorContext(ctx, "failed article upsert rollback", "err", err)
			}
		}
	}()

	if _, err := tx.NamedExecContext(ctx, upsertArticleBodyPostgreSQL, &postgres_dto.ArticleBody{
		ID:                 article.ID,
		ArticleSummariesID: article.ID,
		Body:               article.Body,
		CreatedAt:          article.CreatedAt,
		UpdatedAt:          article.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert article_bodies. err: %w", err)
	}

	if _, err := tx.NamedExecContext(ctx, upsertArticleSummaryPostgreSQL, &postgres_dto.ArticleSummary{
		ID:        article.ID,
		Title:     article.Title,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert article_summaries. err: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed commit. err: %w", err)
	}
	return nil
}