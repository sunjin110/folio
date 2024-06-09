package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type GoogleCustomSearch interface {
	Search(ctx context.Context, searchKeyword string) ([]*model.GoogleSearchResult, error)
}
