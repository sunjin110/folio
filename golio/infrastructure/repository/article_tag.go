package repository

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/postgres"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
	"github.com/sunjin110/folio/golio/infrastructure/repository/query/postgresql"
)

type articleTag struct {
	db *sqlx.DB
}

func NewArticleTag(db *sqlx.DB) repository.ArticleTag {
	return &articleTag{
		db: db,
	}
}

func (a *articleTag) Delete(ctx context.Context, id string) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sql, args := sb.DeleteFrom(postgres.ArticleTagDB).Where(sb.EQ("id", id)).BuildWithFlavor(sqlbuilder.PostgreSQL)
	_, err := a.db.DB.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed delete article tags. id: %s, err: %w", id, err)
	}
	return nil
}

func (a *articleTag) Find(ctx context.Context, sortType repository.SortType, paging *repository.Paging, search *repository.ArticleTagSearch) ([]*model.ArticleTag, error) {
	sb := sqlbuilder.NewStruct(&postgres_dto.ArticleTag{}).SelectFrom(postgres.ArticleTagDB)

	if search != nil && search.Name != nil {
		// TODO PGroona
		sb.Where(sb.Like("name", fmt.Sprintf("%%%s%%", *search.Name)))
	}

	sb.Limit(paging.Limit).Offset(paging.Offset).OrderBy("name")

	if sortType == repository.SortTypeAsc {
		sb.Asc()
	} else {
		sb.Desc()
	}

	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	tags := postgres_dto.ArticleTags{}

	if err := a.db.SelectContext(ctx, &tags, sql, args...); err != nil {
		return nil, fmt.Errorf("failed find article tags. err: %w", err)
	}
	return tags.ToModels(), nil
}

func (a *articleTag) FindByIDs(ctx context.Context, ids []string) ([]*model.ArticleTag, error) {
	if len(ids) == 0 {
		return []*model.ArticleTag{}, nil
	}

	sb := sqlbuilder.NewStruct(&postgres_dto.ArticleTag{}).SelectFrom(postgres.ArticleTagDB)
	sb.Where(sb.In("id", postgres.ToInterfaces(ids)...))
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	tags := postgres_dto.ArticleTags{}

	if err := a.db.SelectContext(ctx, &tags, sql, args...); err != nil {
		return nil, fmt.Errorf("failed find by ids article tags. err: %w", err)
	}
	return tags.ToModels(), nil
}

func (a *articleTag) Get(ctx context.Context, id string) (*model.ArticleTag, error) {
	sb := sqlbuilder.NewStruct(&postgres_dto.ArticleTag{}).SelectFrom(postgres.ArticleTagDB)
	sql, args := sb.Where(sb.Equal("id", id)).BuildWithFlavor(sqlbuilder.PostgreSQL)

	articleTag := &postgres_dto.ArticleTag{}
	if err := a.db.GetContext(ctx, articleTag, sql, args...); err != nil {
		return nil, fmt.Errorf("failed get article tags. id: %s, err: %s", id, err)
	}

	return articleTag.ToModel(), nil
}

func (a *articleTag) Insert(ctx context.Context, tag *model.ArticleTag) error {
	if err := a.upsert(ctx, tag); err != nil {
		return fmt.Errorf("failed insert. err: %w", err)
	}
	return nil
}

func (a *articleTag) Update(ctx context.Context, tag *model.ArticleTag) error {
	if err := a.upsert(ctx, tag); err != nil {
		return fmt.Errorf("failed update. err: %w", err)
	}
	return nil
}

func (a *articleTag) CountTotal(ctx context.Context, search *repository.ArticleTagSearch) (int32, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)").From(postgres.ArticleTagDB)

	if search != nil && search.Name != nil {
		// TODO PGroona
		sb.Where(sb.Like("name", fmt.Sprintf("%%%s%%", *search.Name)))
	}

	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := a.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return -1, fmt.Errorf("failed get countTotal. err: %w", err)
	}

	if !rows.Next() {
		return -1, fmt.Errorf("failed not found rows. sql: %s", sql)
	}

	var totalCount int32
	if err := rows.Scan(&totalCount); err != nil {
		return -1, fmt.Errorf("failed scan. err: %w", err)
	}
	return totalCount, nil
}

func (a *articleTag) upsert(ctx context.Context, tag *model.ArticleTag) error {
	if _, err := a.db.NamedExecContext(ctx, postgresql.UpsertArticleTag, &postgres_dto.ArticleTag{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedTime,
		UpdatedAt: tag.UpdatedTime,
	}); err != nil {
		return fmt.Errorf("failed upsert article_tags. err: %w", err)
	}
	return nil
}
