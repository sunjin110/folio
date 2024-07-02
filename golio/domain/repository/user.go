package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type User interface {
	// Error: ErrNotFound
	Get(ctx context.Context, email string) (*model.User, error)
	Upsert(ctx context.Context, user *model.User) error
}
