package repository

import (
	"context"

	"github.com/sunjin110/folio/lime/domain/model"
)

type Storage interface {
	SaveContent(ctx context.Context, content *model.Content) (err error)
}
