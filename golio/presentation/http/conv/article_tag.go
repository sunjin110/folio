package conv

import (
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func ToArticleTags(tags []*model.ArticleTag) []openapi.ArticleTag {
	openapiTags := make([]openapi.ArticleTag, 0, len(tags))
	for _, tag := range tags {
		openapiTags = append(openapiTags, toArticleTag(tag))
	}
	return openapiTags
}

func toArticleTag(tag *model.ArticleTag) openapi.ArticleTag {
	return openapi.ArticleTag{
		Id:   tag.ID,
		Name: tag.Name,
	}
}
