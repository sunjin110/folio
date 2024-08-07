package repository

import (
	"context"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type GoogleOAuth2 interface {
	GenerateAuthorizationURL() (string, error)
	GetTokenFromCode(ctx context.Context, code string) (*model.Token, error)
	GetTokenFromRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error)
	GetUser(ctx context.Context, token string) (*model.GoogleOAuthUser, error)
	VerifyToken(ctx context.Context, idToken string) (ok bool, expireTime time.Time, err error)
}
