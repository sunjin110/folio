package postgres_dto

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type ArticleBody struct {
	ID                 string    `db:"id"`
	ArticleSummariesID string    `db:"article_summaries_id"`
	Body               string    `db:"body"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type ArticleSummaries []*ArticleSummary

func (summaries ArticleSummaries) ToSummariesModel() []*model.ArticleSummary {
	models := make([]*model.ArticleSummary, 0, len(summaries))
	for _, summary := range summaries {
		models = append(models, summary.toSummaryModel())
	}
	return models
}

type ArticleSummary struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	TagIDs    []string  `db:"tag_ids"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (summary *ArticleSummary) toSummaryModel() *model.ArticleSummary {
	if summary == nil {
		return nil
	}

	return &model.ArticleSummary{
		ID:        summary.ID,
		Title:     summary.Title,
		Writer:    "TODO",
		CreatedAt: summary.CreatedAt,
		UpdatedAt: summary.UpdatedAt,
	}
}
