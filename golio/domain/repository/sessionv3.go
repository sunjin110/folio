package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type SessionV3 interface {
	Upsert(ctx context.Context, userSession *model.UserSessionV3) error
	Get(ctx context.Context, accessToken string) (*model.UserSessionV3, error)
	DeleteByEmail(ctx context.Context, email string) error
}
