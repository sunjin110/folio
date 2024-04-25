package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Authorization interface {
	StartSession(ctx context.Context, token *model.Token, userAuthorization *model.UserAuthorization) error
	Get(ctx context.Context, accessToken string) (*model.UserAuthorization, error)
	CloseSession(ctx context.Context, accessToken string) error
}
