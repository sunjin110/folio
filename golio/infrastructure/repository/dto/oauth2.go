package dto

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type InputGetToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
}

// OutputGetToken https://developers.google.com/identity/protocols/oauth2/web-server?hl=ja#exchange-authorization-code
type OutputGetToken struct {
	AccessToken  string        `json:"access_token"`
	ExpiresIn    time.Duration `json:"expires_in"`    // アクセストークンの残り存続期間(秒)
	RefreshToken string        `json:"refresh_token"` // 新しいアクセストークン
	TokenType    string        `json:"token_type"`
	Scope        string        `json:"scope"`
}

func (o *OutputGetToken) ToModel() *model.Token {
	if o == nil {
		return nil
	}
	return &model.Token{
		AccessToken:  o.AccessToken,
		RefreshToken: o.RefreshToken,
		ExpireTime:   time.Now().Add(o.ExpiresIn * time.Second),
	}
}
