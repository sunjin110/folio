package conv

import (
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
)

func NewArticle(summary *postgres_dto.ArticleSummary, body *postgres_dto.ArticleBody, tagMap map[string]*model.ArticleTag) *model.Article {
	tags := []*model.ArticleTag{}
	for _, tagID := range summary.TagIDs {
		tag, ok := tagMap[tagID]
		if !ok {
			continue
		}
		tags = append(tags, tag)
	}
	return &model.Article{
		ID:        summary.ID,
		Title:     summary.Title,
		Body:      body.Body,
		Writer:    "",
		Tags:      tags,
		CreatedAt: summary.CreatedAt,
		UpdatedAt: summary.UpdatedAt,
	}
}
