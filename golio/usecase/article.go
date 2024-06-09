package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type Article interface {
	Get(ctx context.Context, id string) (*model.Article, error)
	Insert(ctx context.Context, article *model.Article) (*model.Article, error)
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, id string) error
	FindSummaries(ctx context.Context, offset int32, limit int32, titleSearchText *string) (*FindArticleSummariesOutput, error)
	AssistantBodyByAI(ctx context.Context, id string, orderToAI string) (*model.Article, error)
	GenerateArticleByAI(ctx context.Context, prompt string) (*model.Article, error)
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

func (a *article) FindSummaries(ctx context.Context, offset int32, limit int32, titleSearchText *string) (*FindArticleSummariesOutput, error) {
	search := &repository.ArticleSearch{
		Title: titleSearchText,
	}

	summaries, err := a.articleRepo.FindSummary(ctx, repository.SortTypeDesc, &repository.Paging{
		Offset: int(offset),
		Limit:  int(limit),
	}, search)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.FindSummary: %w", err)
	}

	totalCount, err := a.articleRepo.CountTotal(ctx, search)
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

func (a *article) AssistantBodyByAI(ctx context.Context, id string, orderToAI string) (*model.Article, error) {
	article, err := a.articleRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.Get. err: %w", err)
	}

	generatedArticle, err := a.articleRepo.ChangeBodyByAI(ctx, article, orderToAI)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.ChangeBodyByAI. err: %w", err)
	}

	return generatedArticle, nil
}

func (a *article) GenerateArticleByAI(ctx context.Context, prompt string) (*model.Article, error) {
	body, err := a.articleRepo.GenerateBodyByAI(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("fialed articleRepo.GenerateBodyByAI. err: %w", err)
	}

	article := model.NewArticle(prompt, body, "", time.Now())

	if err := a.articleRepo.Insert(ctx, article); err != nil {
		return nil, fmt.Errorf("failed articleRepo.Insert. err: %w", err)
	}
	return article, nil
}
