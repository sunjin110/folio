package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type OAuth2 interface {
	GenerateAuthorizationURL() (string, error)
	GetTokenFromCode(ctx context.Context, code string) (*model.Token, error)
}
