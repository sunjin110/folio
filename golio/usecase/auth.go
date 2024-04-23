package usecase

import (
	"context"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type Auth interface {
	GenerateGoogleAuthorizationURL() (string, error)
	GetGoogleTokenFromCode(ctx context.Context, code string) (*model.Token, error)
	GetUserAuthorizationFromGoogleToken(ctx context.Context, token string) (*model.UserAuthorization, error)
}

type auth struct {
	googleOAuth2 repository.GoogleOAuth2
}

func NewAuth(googleOAuth2 repository.GoogleOAuth2) Auth {
	return &auth{
		googleOAuth2: googleOAuth2,
	}
}

func (a *auth) GenerateGoogleAuthorizationURL() (string, error) {
	url, err := a.googleOAuth2.GenerateAuthorizationURL()
	if err != nil {
		return "", fmt.Errorf("failed googleOAuth2.GenerateAuthorizationURL: %w", err)
	}
	return url, nil
}

func (a *auth) GetGoogleTokenFromCode(ctx context.Context, code string) (*model.Token, error) {
	panic("unimplemented")
}

func (a *auth) GetUserAuthorizationFromGoogleToken(ctx context.Context, token string) (*model.UserAuthorization, error) {
	panic("todo")
}
