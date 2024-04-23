package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
)

const (
	// grantTypeCode codeを利用したやり方
	grantTypeCode = "authorization_code"

	// tokenGetURI トークンを発行するendpointです
	tokenGetURI = "https://oauth2.googleapis.com/token"
)

type oauth2 struct {
	clientID     string
	clientSecret string
	redirectURI  string
}

func NewOAuth2(ctx context.Context, clientID string, clientSecret string, redirectURI string) (repository.OAuth2, error) {

	return &oauth2{
		clientID:     clientID,
		clientSecret: clientID,
		redirectURI:  redirectURI,
	}, nil
}

// InputGetToken tokenを取得するために必要なinput
type InputGetToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`   // authorization_code
	RedirectURI  string `json:"redirect_uri"` // http://localhost:3001/auth/google-oauth/callback
}

// GenerateAuthorizationURL Clientが叩くべき認証のURLを作成する
func (o *oauth2) GenerateAuthorizationURL() (string, error) {
	u, err := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")
	if err != nil {
		return "", fmt.Errorf("failed url.Parse: %w", err)
	}

	q := u.Query()
	q.Set("client_id", o.clientID)
	q.Set("redirect_uri", o.redirectURI)
	q.Set("response_type", "code")
	q.Set("scope", "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")
	q.Set("access_type", "offline")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// GetTokenFromCode　取得したcodeからTokenを取得する
func (o *oauth2) GetTokenFromCode(ctx context.Context, code string) (*model.Token, error) {
	formData := url.Values{}
	formData.Set("client_id", o.clientID)
	formData.Set("client_secret", o.clientSecret)
	formData.Set("code", code)
	formData.Set("redirect_uri", o.redirectURI)
	formData.Set("grant_type", grantTypeCode)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokenGetURI, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client.Do: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll: %w", err)
	}

	output := &dto.OutputGetToken{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal: body: %s, err: %w", string(body), err)
	}
	return output.ToModel(), nil
}
