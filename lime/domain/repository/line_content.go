package repository

import (
	"context"
	"io"

	"github.com/sunjin110/folio/lime/domain/model"
)

type LineContent interface {
	GetByProvider(ctx context.Context, provider model.LineContentProvider) (io.ReadCloser, error)
}
