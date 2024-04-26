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
	GetUserSessionFromGoogleToken(ctx context.Context, token string) (*model.UserSession, error)

	// StartSessionFromGoogleOAuthCode google oauth2.0のcodeからsessionを開始する
	StartSessionFromGoogleOAuthCode(ctx context.Context, code string) (*StartSessionOutput, error)
}

type StartSessionOutput struct {
	AccessToken string
	Email       string
}

type auth struct {
	googleOAuth2 repository.GoogleOAuth2
	sessionRepo  repository.Session
}

func NewAuth(googleOAuth2 repository.GoogleOAuth2, sessionRepo repository.Session) Auth {
	return &auth{
		googleOAuth2: googleOAuth2,
		sessionRepo:  sessionRepo,
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

func (a *auth) GetUserSessionFromGoogleToken(ctx context.Context, token string) (*model.UserSession, error) {

	panic("todo")
}

func (a *auth) StartSessionFromGoogleOAuthCode(ctx context.Context, code string) (*StartSessionOutput, error) {

	token, err := a.googleOAuth2.GetTokenFromCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed googleOAuth2.GetTokenFromCode: %w", err)
	}

	userSession, err := a.googleOAuth2.GetUserSession(ctx, token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed googleOAuth2.GetUserSession: %w", err)
	}

	if err := a.sessionRepo.Start(ctx, token, userSession); err != nil {
		return nil, fmt.Errorf("failed sessionRepo.StartSession: %w", err)
	}

	return &StartSessionOutput{
		AccessToken: token.AccessToken,
		Email:       userSession.Email,
	}, nil
}
