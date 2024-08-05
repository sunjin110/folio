package dto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type InputGetGoogleToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
}

// OutputGetGoogleToken https://developers.google.com/identity/protocols/oauth2/web-server?hl=ja#exchange-authorization-code
type OutputGetGoogleToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int32  `json:"expires_in"`    // アクセストークンの残り存続期間(秒)
	RefreshToken string `json:"refresh_token"` // 新しいアクセストークン
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

func (o *OutputGetGoogleToken) ToModel() *model.Token {
	if o == nil {
		return nil
	}
	return &model.Token{
		AccessToken:  o.AccessToken,
		RefreshToken: o.RefreshToken,
		ExpireTime:   time.Now().Add(time.Duration(o.ExpiresIn) * time.Second),
	}
}

// https://developers.google.com/identity/protocols/oauth2/web-server?hl=ja#offline
type InputRefreshGoogleToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`    // refresh_token
	RefreshToken string `json:"refresh_token"` // 認証コード返還から返された更新トークン
}

type OutputRefreshGoogleToken struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int32  `json:"expires_in"` // アクセストークンの残り時間(秒)
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (o *OutputRefreshGoogleToken) ToToken(refreshToken string) *model.Token {
	if o == nil {
		return nil
	}
	return &model.Token{
		AccessToken:  o.AccessToken,
		RefreshToken: refreshToken,
		ExpireTime:   time.Now().Add(time.Duration(o.ExpireIn) * time.Second),
	}
}

// Tokeninfoの内容 https://cloud.google.com/docs/authentication/token-types?hl=ja#id-contents
type TokenInfo struct {
	Iss string `json:"iss"`
	Azp string `json:"azp"`
	Aud string `json:"aud"`
	Sub string `json:"sub"`
	Iat string `json:"iat"`
	Exp string `json:"exp"` // Tokenが無くなる日
}

func (tokeninfo *TokenInfo) GetExpireTime() (time.Time, error) {
	expEpoch, err := strconv.Atoi(tokeninfo.Exp)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed atoi exp. exp: %s, err: %w", tokeninfo.Exp, err)
	}
	return time.Unix(int64(expEpoch), 0), nil
}
