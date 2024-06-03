package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type GoogleOAuth2 interface {
	GenerateAuthorizationURL() (string, error)
	GetTokenFromCode(ctx context.Context, code string) (*model.Token, error)
	GetTokenFromRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error)
	GetUserSession(ctx context.Context, token string) (*model.UserSession, error)
}
