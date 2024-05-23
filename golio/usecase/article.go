package usecase

import (
	"context"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type Article interface {
	Get(ctx context.Context, id string) (*model.Article, error)
	Insert(ctx context.Context, article *model.Article) (*model.Article, error)
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, id string) error
	FindSummaries(ctx context.Context, offset int32, limit int32) (*FindArticleSummariesOutput, error)
}

type FindArticleSummariesOutput struct {
	Articles   []*model.ArticleSummary
	TotalCount int32
}

type article struct {
	articleRepo repository.Article
}

func NewArticle(articleRepo repository.Article) Article {
	return &article{
		articleRepo: articleRepo,
	}
}

func (a *article) Delete(ctx context.Context, id string) error {
	if err := a.articleRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed articleRepo.Delete: %w", err)
	}
	return nil
}

func (a *article) FindSummaries(ctx context.Context, offset int32, limit int32) (*FindArticleSummariesOutput, error) {
	summaries, err := a.articleRepo.FindSummary(ctx, repository.SortTypeDesc, &repository.Paging{
		Offset: int(offset),
		Limit:  int(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.FindSummary: %w", err)
	}

	totalCount, err := a.articleRepo.CountTotal(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.CountTotal: %w", err)
	}

	return &FindArticleSummariesOutput{
		Articles:   summaries,
		TotalCount: totalCount,
	}, nil
}

func (a *article) Get(ctx context.Context, id string) (*model.Article, error) {
	article, err := a.articleRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.Get: %w", err)
	}
	return article, nil
}

func (a *article) Insert(ctx context.Context, article *model.Article) (*model.Article, error) {
	if err := a.articleRepo.Insert(ctx, article); err != nil {
		return nil, fmt.Errorf("failed articleRepo.Insert: %w", err)
	}
	return article, nil
}

func (a *article) Update(ctx context.Context, article *model.Article) error {
	if err := a.articleRepo.Update(ctx, article); err != nil {
		return fmt.Errorf("failed articleRepo.Update: %w", err)
	}
	return nil
}
