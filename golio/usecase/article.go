package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/usecase/input"
)

type Article interface {
	Get(ctx context.Context, id string) (*model.Article, error)
	Insert(ctx context.Context, input *input.ArticleInsert) (*model.Article, error)
	Update(ctx context.Context, input *input.ArticleUpdate) error
	Delete(ctx context.Context, id string) error
	FindSummaries(ctx context.Context, offset int32, limit int32, titleSearchText *string, tags []string) (*FindArticleSummariesOutput, error)
	AssistantBodyByAI(ctx context.Context, id string, orderToAI string) (*model.Article, error)
	GenerateArticleByAI(ctx context.Context, prompt string) (*model.Article, error)
	FindTags(ctx context.Context, offset int32, limit int32, nameSearchText *string) ([]*model.ArticleTag, error)
	InsertTag(ctx context.Context, tag *model.ArticleTag) error
	UpdateTag(ctx context.Context, tag *model.ArticleTag) error
	DeleteTag(ctx context.Context, tagID string) error
}

type FindArticleSummariesOutput struct {
	Articles   []*model.ArticleSummary
	TotalCount int32
}

type article struct {
	articleRepo    repository.Article
	articleAIRepo  repository.ArticleAI
	articleTagRepo repository.ArticleTag
}

func NewArticle(articleRepo repository.Article, articleAIRepo repository.ArticleAI, articleTagRepo repository.ArticleTag) Article {
	return &article{
		articleRepo:    articleRepo,
		articleAIRepo:  articleAIRepo,
		articleTagRepo: articleTagRepo,
	}
}

func (a *article) Delete(ctx context.Context, id string) error {
	if err := a.articleRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed articleRepo.Delete: %w", err)
	}
	return nil
}

func (a *article) FindSummaries(ctx context.Context, offset int32, limit int32, titleSearchText *string, tags []string) (*FindArticleSummariesOutput, error) {
	search := &repository.ArticleSearch{
		Title: titleSearchText,
		Tags:  tags,
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

func (a *article) Insert(ctx context.Context, input *input.ArticleInsert) (*model.Article, error) {
	tags, err := a.articleTagRepo.FindByIDs(ctx, input.TagIDs)
	if err != nil {
		return nil, fmt.Errorf("failed articleTagRepo.FindByIDs. err: %w", err)
	}

	article := model.NewArticle(input.Title, input.Body, "", tags, time.Now())
	if err := a.articleRepo.Insert(ctx, article); err != nil {
		return nil, fmt.Errorf("failed articleRepo.Insert: %w", err)
	}
	return article, nil
}

func (a *article) Update(ctx context.Context, input *input.ArticleUpdate) error {
	tags, err := a.articleTagRepo.FindByIDs(ctx, input.TagIDs)
	if err != nil {
		return fmt.Errorf("failed articleTagRepo.FindByIDs. err: %w", err)
	}

	beforeArticle, err := a.articleRepo.Get(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("failed get article. err: %w", err)
	}

	article := &model.Article{
		ID:        input.ID,
		Title:     input.Title,
		Body:      input.Body,
		Writer:    "",
		Tags:      tags,
		CreatedAt: beforeArticle.CreatedAt,
		UpdatedAt: beforeArticle.UpdatedAt,
	}

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

	generatedArticle, err := a.articleAIRepo.ChangeBodyByAI(ctx, article, orderToAI)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.ChangeBodyByAI. err: %w", err)
	}

	return generatedArticle, nil
}

func (a *article) GenerateArticleByAI(ctx context.Context, prompt string) (*model.Article, error) {
	body, err := a.articleAIRepo.GenerateBodyByAI(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed articleRepo.GenerateBodyByAI. err: %w", err)
	}

	article := model.NewArticle(prompt, body, "", nil, time.Now())

	if err := a.articleRepo.Insert(ctx, article); err != nil {
		return nil, fmt.Errorf("failed articleRepo.Insert. err: %w", err)
	}
	return article, nil
}

func (a *article) FindTags(ctx context.Context, offset int32, limit int32, nameSearchText *string) ([]*model.ArticleTag, error) {
	tags, err := a.articleTagRepo.Find(ctx, repository.SortTypeAsc, &repository.Paging{
		Offset: int(offset),
		Limit:  int(limit),
	}, &repository.ArticleTagSearch{
		Name: nameSearchText,
	})
	if err != nil {
		return nil, fmt.Errorf("failed articleTagRepo.Find. err: %w", err)
	}
	return tags, nil
}

func (a *article) InsertTag(ctx context.Context, tag *model.ArticleTag) error {
	if err := a.articleTagRepo.Insert(ctx, tag); err != nil {
		return fmt.Errorf("failed articleTagRepo.Insert. err: %w", err)
	}
	return nil
}

func (a *article) UpdateTag(ctx context.Context, tag *model.ArticleTag) error {
	if err := a.articleTagRepo.Update(ctx, tag); err != nil {
		return fmt.Errorf("failed articleTagRepo.Update. err: %w", err)
	}
	return nil
}

func (a *article) DeleteTag(ctx context.Context, tagID string) error {
	if err := a.articleTagRepo.Delete(ctx, tagID); err != nil {
		return fmt.Errorf("failed articleTagRepo.Delete. err: %w", err)
	}
	return nil
}
