package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/gdto"
	"google.golang.org/api/oauth2/v2"
)

const (
	// grantTypeCode codeを利用したやり方
	grantTypeCode = "authorization_code"

	// refreshGrantType refresh tokenを利用したやり方
	refreshGrantType = "refresh_token"

	// tokenGetURI トークンを発行するendpointです
	tokenGetURI = "https://oauth2.googleapis.com/token"

	// personGetURI ユーザーの情報を取得するendpoint
	personGetURI = "https://people.googleapis.com/v1/people/me"
)

type googleOauth2 struct {
	clientID      string
	clientSecret  string
	redirectURI   string
	oauth2Service *oauth2.Service
}

func NewGoogleOAuth2(ctx context.Context, clientID string, clientSecret string, redirectURI string) (repository.GoogleOAuth2, error) {
	oauth2Service, err := oauth2.NewService(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed oauth2.NewService. err: %w", err)
	}

	return &googleOauth2{
		clientID:      clientID,
		clientSecret:  clientSecret,
		redirectURI:   redirectURI,
		oauth2Service: oauth2Service,
	}, nil
}

// GenerateAuthorizationURL Clientが叩くべき認証のURLを作成する
func (o *googleOauth2) GenerateAuthorizationURL() (string, error) {
	u, err := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")
	if err != nil {
		return "", fmt.Errorf("failed url.Parse: %w", err)
	}

	q := u.Query()
	q.Set("client_id", o.clientID)
	q.Set("redirect_uri", o.redirectURI)
	q.Set("response_type", "code")
	q.Set("scope", "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")
	q.Set("access_type", "offline") // prompt=consent
	q.Set("prompt", "consent")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// GetTokenFromCode　取得したcodeからTokenを取得する
func (o *googleOauth2) GetTokenFromCode(ctx context.Context, code string) (*model.Token, error) {
	reqBody := &dto.InputGetGoogleToken{
		ClientID:     o.clientID,
		ClientSecret: o.clientSecret,
		Code:         code,
		RedirectURI:  o.redirectURI,
		GrantType:    grantTypeCode,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed json.Marshal: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokenGetURI, strings.NewReader(string(b)))
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client.Do: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed GetTokenFromCode request. statusCode: %d, body: %s", resp.StatusCode, string(body))
	}

	output := &dto.OutputGetGoogleToken{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal: body: %s, err: %w", string(body), err)
	}
	return output.ToModel(), nil
}

func (o *googleOauth2) GetTokenFromRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error) {
	reqBody := &dto.InputRefreshGoogleToken{
		ClientID:     o.clientID,
		ClientSecret: o.clientSecret,
		GrantType:    refreshGrantType,
		RefreshToken: refreshToken,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed json.Marshal. err: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokenGetURI, strings.NewReader(string(b)))
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest. err: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client.Do. err: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll. err: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed GetTokenFromRefreshToken request. statusCode: %d, body: %s", resp.StatusCode, string(body))
	}

	output := &dto.OutputRefreshGoogleToken{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. body: %s, err: %w", string(body), err)
	}
	return output.ToToken(refreshToken), nil
}

func (o *googleOauth2) GetUser(ctx context.Context, token string) (*model.GoogleOAuthUser, error) {
	u, err := url.Parse(personGetURI)
	if err != nil {
		return nil, fmt.Errorf("failed url.Parse: %w", err)
	}

	q := u.Query()
	q.Set("personFields", "names,emailAddresses")
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client.Do: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll: %w", err)
	}

	person := &gdto.Person{}
	if err := json.Unmarshal(body, person); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. body: %s, err: %w", string(body), err)
	}
	primaryName := person.GetPrimaryName()
	return &model.GoogleOAuthUser{
		Email:       person.GetPrimaryEmailAddress().Value,
		FirstName:   primaryName.GivenName,
		LastName:    primaryName.FamilyName,
		DisplayName: primaryName.DisplayName,
	}, nil
}

func (o *googleOauth2) VerifyToken(ctx context.Context, token string, accessToken *string) (bool, time.Time, error) {
	call := o.oauth2Service.Tokeninfo()
	call.Context(ctx)
	call.IdToken(token)
	if accessToken != nil {
		call.AccessToken(*accessToken)
	}

	tokeninfo, err := call.Do()
	if err != nil {
		return false, time.Time{}, fmt.Errorf("failed call tokeninfo. err: %w", err)
	}

	return true, time.Now().Add(time.Duration(tokeninfo.ExpiresIn) * time.Second), nil
}
