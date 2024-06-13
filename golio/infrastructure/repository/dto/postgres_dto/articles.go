package postgres_dto

import (
	"time"

	"github.com/lib/pq"
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

func (summaries ArticleSummaries) ToSummariesModel(tagMap map[string]*model.ArticleTag) []*model.ArticleSummary {
	models := make([]*model.ArticleSummary, 0, len(summaries))
	for _, summary := range summaries {
		models = append(models, summary.toSummaryModel(tagMap))
	}
	return models
}

func (summaries ArticleSummaries) GetTagIDs() []string {
	tagMap := map[string]struct{}{}
	for _, summary := range summaries {
		for _, tag := range summary.TagIDs {
			tagMap[tag] = struct{}{}
		}
	}

	tags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		tags = append(tags, tag)
	}
	return tags
}

type ArticleSummary struct {
	ID        string         `db:"id"`
	Title     string         `db:"title"`
	TagIDs    pq.StringArray `db:"tag_ids"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

func (summary *ArticleSummary) toSummaryModel(tagMap map[string]*model.ArticleTag) *model.ArticleSummary {
	if summary == nil {
		return nil
	}

	tags := make([]*model.ArticleTag, 0, len(summary.TagIDs))
	for _, tagID := range summary.TagIDs {
		if tag, ok := tagMap[tagID]; ok {
			tags = append(tags, tag)
		}
	}

	return &model.ArticleSummary{
		ID:        summary.ID,
		Title:     summary.Title,
		Writer:    "TODO",
		Tags:      tags,
		CreatedAt: summary.CreatedAt,
		UpdatedAt: summary.UpdatedAt,
	}
}
