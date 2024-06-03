package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type Auth interface {
	// GenerateGoogleAuthorizationURL google oauth2.0のredirectをURLを返す
	GenerateGoogleAuthorizationURL() (string, error)

	// StartSessionFromGoogleOAuthCode google oauth2.0のcodeからsessionを開始する
	StartSessionFromGoogleOAuthCode(ctx context.Context, code string) (*StartSessionOutput, error)

	// GetSessionInfoFromToken Session情報をtokenから取得する
	GetSessionInfoFromToken(ctx context.Context, token string) (*model.UserSessionV2, error)

	// RefreshSessionFromRefreshToken Session情報をrefreshTokenを利用して綺麗にする
	RefreshSessionFromRefreshToken(ctx context.Context, userSession *model.UserSessionV2) (*model.UserSessionV2, error)
}

type StartSessionOutput struct {
	AccessToken string
	Email       string
}

type auth struct {
	googleOAuth2  repository.GoogleOAuth2
	sessionV2Repo repository.SessionV2
}

func NewAuth(googleOAuth2 repository.GoogleOAuth2, sessionRepoV2 repository.SessionV2) Auth {
	return &auth{
		googleOAuth2:  googleOAuth2,
		sessionV2Repo: sessionRepoV2,
	}
}

func (a *auth) GenerateGoogleAuthorizationURL() (string, error) {
	url, err := a.googleOAuth2.GenerateAuthorizationURL()
	if err != nil {
		return "", fmt.Errorf("failed googleOAuth2.GenerateAuthorizationURL: %w", err)
	}
	return url, nil
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

	if err := a.sessionV2Repo.Upsert(ctx, &model.UserSessionV2{
		Email:                 userSession.Email,
		FirstName:             userSession.FirstName,
		LastName:              userSession.LastName,
		DisplayName:           userSession.DisplayName,
		AccessToken:           token.AccessToken,
		RefreshToken:          token.RefreshToken,
		AccessTokenExpireTime: token.ExpireTime,
	}); err != nil {
		return nil, fmt.Errorf("failed sessionV2Repo.Upsert. err: %w", err)
	}

	return &StartSessionOutput{
		AccessToken: token.AccessToken,
		Email:       userSession.Email,
	}, nil
}

func (a *auth) GetSessionInfoFromToken(ctx context.Context, token string) (*model.UserSessionV2, error) {
	userSessionV2, err := a.sessionV2Repo.GetByAccessToken(ctx, token)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed get sessionV2. err: %w", err)
	}
	return userSessionV2, nil
}

func (a *auth) RefreshSessionFromRefreshToken(ctx context.Context, userSession *model.UserSessionV2) (*model.UserSessionV2, error) {
	token, err := a.googleOAuth2.GetTokenFromRefreshToken(ctx, userSession.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed get token from refresh token. err: %w", err)
	}

	refreshedUserSession := &model.UserSessionV2{
		Email:                 userSession.Email,
		FirstName:             userSession.FirstName,
		LastName:              userSession.LastName,
		DisplayName:           userSession.DisplayName,
		AccessToken:           token.AccessToken,
		RefreshToken:          userSession.RefreshToken,
		AccessTokenExpireTime: token.ExpireTime,
	}

	if err := a.sessionV2Repo.Upsert(ctx, refreshedUserSession); err != nil {
		return nil, fmt.Errorf("failed sessionV2Repo.Upsert. err: %w", err)
	}
	return refreshedUserSession, nil
}
