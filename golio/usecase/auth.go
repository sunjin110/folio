package usecase

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Auth interface {
	GenerateGoogleAuthorizationURL() (string, error)
	GetGoogleTokenFromCode(ctx context.Context, code string) (*model.Token, error)
	GetUserAuthorizationFromGoogleToken(ctx context.Context, token string) (*model.UserAuthorization, error)
}

type auth struct {
}

func NewAuth() Auth {
	return &auth{}
}

func (a *auth) GenerateGoogleAuthorizationURL() (string, error) {
	panic("unimplemented")
}

func (a *auth) GetGoogleTokenFromCode(ctx context.Context, code string) (*model.Token, error) {
	panic("unimplemented")
}

func (a *auth) GetUserAuthorizationFromGoogleToken(ctx context.Context, token string) (*model.UserAuthorization, error) {
	panic("todo")
}
