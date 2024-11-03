package repository

import (
	"context"

	"github.com/sunjin110/folio/gomb/domain/model"
)

type Storage interface {
	SaveContent(ctx context.Context, dir string, content *model.Content) error
	GetContent(ctx context.Context, path string) (*model.Content, error)
}
