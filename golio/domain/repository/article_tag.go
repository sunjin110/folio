package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

//go:generate mockgen -source ./article_tag.go -destination ./mock_repo/article_tag.go -package mock_repo
type ArticleTag interface {
	Get(ctx context.Context, id string) (*model.ArticleTag, error)
	Insert(ctx context.Context, tag *model.ArticleTag) error
	Update(ctx context.Context, tag *model.ArticleTag) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, sortType SortType, paging *Paging, search *ArticleTagSearch) ([]*model.ArticleTag, error)
	FindByIDs(ctx context.Context, ids []string) ([]*model.ArticleTag, error)
	CountTotal(ctx context.Context, search *ArticleTagSearch) (int32, error)
}

type ArticleTagSearch struct {
	Name *string
}
