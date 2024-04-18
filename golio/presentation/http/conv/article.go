package conv

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func ToArticlesGet(articles []*model.ArticleSummary, totalCount int32) openapi.ArticlesGet200Response {
	return openapi.ArticlesGet200Response{
		Articles: toArticlesGet200ResponseArticlesInners(articles),
		Total:    totalCount,
	}
}

func toArticlesGet200ResponseArticlesInners(articles []*model.ArticleSummary) []openapi.ArticlesGet200ResponseArticlesInner {
	inners := make([]openapi.ArticlesGet200ResponseArticlesInner, 0, len(articles))
	for _, article := range articles {
		inners = append(inners, toArticlesGet200ResponseArticlesInner(article))
	}
	return inners
}

func toArticlesGet200ResponseArticlesInner(article *model.ArticleSummary) openapi.ArticlesGet200ResponseArticlesInner {
	return openapi.ArticlesGet200ResponseArticlesInner{
		Id:        article.ID,
		Title:     article.Title,
		CreatedAt: article.CreatedAt.Format(time.RFC3339),
	}
}
