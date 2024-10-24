package repository

import (
	"context"

	"github.com/sunjin110/folio/lime/domain/model"
)

type LineContent interface {
	GetByProvider(ctx context.Context, provider model.LineContentProvider) (*model.Content, error)
}
