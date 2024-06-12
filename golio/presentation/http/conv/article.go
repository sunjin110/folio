package conv

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func ToArticlesGet(articles []*model.ArticleSummary, totalCount int32) openapi.ArticlesGet200Response {
	return openapi.ArticlesGet200Response{
		Articles: toArticles(articles),
		Total:    totalCount,
	}
}

func toArticles(articles []*model.ArticleSummary) []openapi.Article {
	inners := make([]openapi.Article, 0, len(articles))
	for _, article := range articles {
		inners = append(inners, toArticle(article))
	}
	return inners
}

func toArticle(article *model.ArticleSummary) openapi.Article {
	return openapi.Article{
		Id:        article.ID,
		Title:     article.Title,
		CreatedAt: article.CreatedAt.Format(time.RFC3339),
		Tags:      ToArticleTags(article.Tags),
	}
}
