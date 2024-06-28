package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type User interface {
	Get(ctx context.Context, email string) (*model.User, error)
	Upsert(ctx context.Context, user *model.User) error
}
