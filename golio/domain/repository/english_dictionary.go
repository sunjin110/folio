package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type EnglishDictionary interface {
	// GetDetail
	// Error: ErrNotFound
	GetDetail(ctx context.Context, englishWord string) (*model.WordDetail, error)
}
