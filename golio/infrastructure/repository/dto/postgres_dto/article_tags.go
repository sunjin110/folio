package postgres_dto

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type ArticleTags []*ArticleTag

func (tags ArticleTags) ToModels() []*model.ArticleTag {
	models := make([]*model.ArticleTag, 0, len(tags))
	for _, tag := range tags {
		models = append(models, tag.ToModel())
	}
	return models
}

type ArticleTag struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (tag *ArticleTag) ToModel() *model.ArticleTag {
	return &model.ArticleTag{
		ID:          tag.ID,
		Name:        tag.Name,
		CreatedTime: tag.CreatedAt,
		UpdatedTime: tag.UpdatedAt,
	}
}
