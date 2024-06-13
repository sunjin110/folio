package repository

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	_ "embed"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/conv"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
	"github.com/sunjin110/folio/golio/infrastructure/repository/query/postgresql"
	"github.com/sunjin110/folio/golio/utils"
	"golang.org/x/sync/errgroup"
)

type articleV2 struct {
	db             *sqlx.DB
	articleTagRepo repository.ArticleTag
}

func NewArticleV2(ctx context.Context, db *sqlx.DB, articleTagRepo repository.ArticleTag) repository.Article {
	return &articleV2{
		db:             db,
		articleTagRepo: articleTagRepo,
	}
}

func (a *articleV2) CountTotal(ctx context.Context, search *repository.ArticleSearch) (int32, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)").From("article_summaries")

	if search != nil && search.Title != nil {
		// TODO indexが利用できないため将来的にPGroongaを利用する
		sb.Where(sb.Like("title", "%"+*search.Title+"%"))
	}

	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := a.db.QueryxContext(ctx, sql, args...)
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

func (a *articleV2) FindSummary(ctx context.Context, sortType repository.SortType, paging *repository.Paging, search *repository.ArticleSearch) ([]*model.ArticleSummary, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("article_summaries")

	if search != nil {
		if search.Title != nil {
			// TODO indexが利用できないため将来的にPGroongaを利用する
			sb.Where(sb.Like("title", "%"+*search.Title+"%"))
		}

		if len(search.Tags) > 0 {
			args := make([]string, 0, len(search.Tags))
			for _, tag := range search.Tags {
				args = append(args, sb.Cond.Args.Add(tag))
			}
			sb.Where(fmt.Sprintf("tag_ids @> array[%s]", strings.Join(args, ",")))
		}
	}

	sb.Limit(paging.Limit).Offset(paging.Offset).OrderBy("created_at")
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

	tags, err := a.articleTagRepo.FindByIDs(ctx, summaries.GetTagIDs())
	if err != nil {
		return nil, fmt.Errorf("failed get article tags. err: %w", err)
	}

	tagMap := utils.SliceToMap(tags, func(v *model.ArticleTag) string {
		return v.ID
	})
	return summaries.ToSummariesModel(tagMap), nil
}

func (a *articleV2) Get(ctx context.Context, id string) (*model.Article, error) {
	eg, ctx := errgroup.WithContext(ctx)

	body := &postgres_dto.ArticleBody{}
	eg.Go(func() error {
		if err := a.db.GetContext(ctx, body, "select * from article_bodies where id = $1", id); err != nil {
			return fmt.Errorf("failed get article_bodies. err: %w", err)
		}
		return nil
	})

	summary := &postgres_dto.ArticleSummary{}
	tagMap := map[string]*model.ArticleTag{}
	eg.Go(func() error {
		if err := a.db.GetContext(ctx, summary, "select * from article_summaries where id = $1", id); err != nil {
			return fmt.Errorf("failed get article_summaries. err: %w", err)
		}

		if len(summary.TagIDs) == 0 {
			return nil
		}

		tags, err := a.articleTagRepo.FindByIDs(ctx, summary.TagIDs)
		if err != nil {
			return nil
		}

		m := map[string]*model.ArticleTag{}
		for _, tag := range tags {
			m[tag.ID] = tag
		}
		tagMap = m
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed get article. err: %w", err)
	}
	return conv.NewArticle(summary, body, tagMap), nil
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

	if _, err := tx.NamedExecContext(ctx, postgresql.UpsertArticleBody, &postgres_dto.ArticleBody{
		ID:                 article.ID,
		ArticleSummariesID: article.ID,
		Body:               article.Body,
		CreatedAt:          article.CreatedAt,
		UpdatedAt:          article.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert article_bodies. err: %w", err)
	}

	if _, err := tx.NamedExecContext(ctx, postgresql.UpsertArticleSummary, &postgres_dto.ArticleSummary{
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
