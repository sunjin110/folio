package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Session interface {
	Start(ctx context.Context, token *model.Token, userSession *model.UserSession) error
	Get(ctx context.Context, accessToken string) (*model.UserSession, error)
	Close(ctx context.Context, accessToken string) error
}
