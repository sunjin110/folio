package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type SessionV2 interface {
	Upsert(ctx context.Context, userSession *model.UserSessionV2) error
	GetByAccessToken(ctx context.Context, accessToken string) (*model.UserSessionV2, error)
	DeleteByAccessToken(ctx context.Context, accessToken string) error
}
