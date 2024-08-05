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
	GetSessionInfoFromToken(ctx context.Context, token string) (*model.UserSessionV3, error)

	RefreshSession(ctx context.Context, refreshToken string, email string) (*model.UserSessionV3, error)

	VerifyTokenAndStartSession(ctx context.Context, token string, accessToken string, refreshToken string) (*StartSessionOutput, error)
}

type StartSessionOutput struct {
	AccessToken string
	Email       string
}

type auth struct {
	googleOAuth2  repository.GoogleOAuth2
	userRepo      repository.User
	sessionV3Repo repository.SessionV3
}

func NewAuth(googleOAuth2 repository.GoogleOAuth2, userRepo repository.User, sessionRepoV3 repository.SessionV3) Auth {
	return &auth{
		googleOAuth2:  googleOAuth2,
		userRepo:      userRepo,
		sessionV3Repo: sessionRepoV3,
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

	googleOAuthUser, err := a.googleOAuth2.GetUser(ctx, token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed googleOAuth2.GetUserSession: %w", err)
	}

	if err := a.userRepo.Upsert(ctx, &model.User{
		Email:        googleOAuthUser.Email,
		RefreshToken: token.RefreshToken,
		FirstName:    googleOAuthUser.FirstName,
		LastName:     googleOAuthUser.LastName,
		DisplayName:  googleOAuthUser.DisplayName,
	}); err != nil {
		return nil, fmt.Errorf("failed userRepo.Upsert. err: %w", err)
	}

	if err := a.sessionV3Repo.Upsert(ctx, &model.UserSessionV3{
		AccessToken:           token.AccessToken,
		Email:                 googleOAuthUser.Email,
		AccessTokenExpireTime: token.ExpireTime,
	}); err != nil {
		return nil, fmt.Errorf("failed sessionV3Repo.Upsert. err: %w", err)
	}

	return &StartSessionOutput{
		AccessToken: token.AccessToken,
		Email:       googleOAuthUser.Email,
	}, nil
}

func (a *auth) GetSessionInfoFromToken(ctx context.Context, token string) (*model.UserSessionV3, error) {
	userSession, err := a.sessionV3Repo.Get(ctx, token)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed get userSessionV3. err: %w", err)
	}

	return userSession, nil
}

func (a *auth) RefreshSession(ctx context.Context, refreshToken string, email string) (*model.UserSessionV3, error) {
	token, err := a.googleOAuth2.GetTokenFromRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed get token from refresh token. err: %w", err)
	}

	if err := a.sessionV3Repo.Upsert(ctx, &model.UserSessionV3{
		AccessToken:           token.AccessToken,
		Email:                 email,
		AccessTokenExpireTime: token.ExpireTime,
	}); err != nil {
		return nil, fmt.Errorf("failed upsert user session. err: %w", err)
	}

	return &model.UserSessionV3{
		AccessToken:           token.AccessToken,
		Email:                 email,
		AccessTokenExpireTime: token.ExpireTime,
	}, nil
}

func (a *auth) VerifyTokenAndStartSession(ctx context.Context, token string, accessToken string, refreshToken string) (*StartSessionOutput, error) {

	ok, exipireTime, err := a.googleOAuth2.VerifyToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed googleOAuth2.VerifyToken. err: %w", err)
	}
	if !ok {
		return nil, ErrPermissionDenied
	}

	googleOAuthUser, err := a.googleOAuth2.GetUser(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed googleOAuth2.GetUserSession. err: %w", err)
	}

	if err := a.userRepo.Upsert(ctx, &model.User{
		Email:        googleOAuthUser.Email,
		RefreshToken: refreshToken,
		FirstName:    googleOAuthUser.FirstName,
		LastName:     googleOAuthUser.LastName,
		DisplayName:  googleOAuthUser.DisplayName,
	}); err != nil {
		return nil, fmt.Errorf("failed userRepo.Upsert. err: %w", err)
	}

	if err := a.sessionV3Repo.Upsert(ctx, &model.UserSessionV3{
		AccessToken:           accessToken,
		Email:                 googleOAuthUser.Email,
		AccessTokenExpireTime: exipireTime,
	}); err != nil {
		return nil, fmt.Errorf("failed sessionV3Repo.Upsert. err: %w", err)
	}

	return &StartSessionOutput{
		AccessToken: accessToken,
		Email:       googleOAuthUser.Email,
	}, nil

}
