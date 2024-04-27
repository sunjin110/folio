package conv

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
)

func ToArticleSummaries(output *d1.Output) []*model.ArticleSummary {
	if output == nil {
		return nil
	}
	summaries := output.GetResultMapList()
	models := make([]*model.ArticleSummary, 0, len(summaries))
	for _, summary := range summaries {
		if summaryModel := toArticleSummary(summary); summaryModel != nil {
			models = append(models, summaryModel)
		}
	}
	return models
}

func toArticleSummary(m map[string]interface{}) *model.ArticleSummary {
	if len(m) == 0 {
		return nil
	}
	return &model.ArticleSummary{
		ID:        m["id"].(string),
		Title:     m["title"].(string),
		Writer:    "todo",
		CreatedAt: time.Unix(int64(m["created_at"].(float64)), 0),
		UpdatedAt: time.Unix(int64(m["updated_at"].(float64)), 0),
	}
}
