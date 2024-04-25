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
	authRepo     repository.Authorization
}

func NewAuth(googleOAuth2 repository.GoogleOAuth2, authRepo repository.Authorization) Auth {
	return &auth{
		googleOAuth2: googleOAuth2,
		authRepo:     authRepo,
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
	token, err := a.googleOAuth2.GetTokenFromCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed googleOAuth2.GetTokenFromCode: %w", err)
	}
	return token, nil
}

func (a *auth) GetUserAuthorizationFromGoogleToken(ctx context.Context, token string) (*model.UserAuthorization, error) {

	panic("todo")
}
