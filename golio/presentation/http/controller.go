package http

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/sunjin110/folio/golio/usecase"
)

type googleOAuthController struct {
	authUsecase         usecase.Auth
	callbackRedirectURI string
}

func NewGoogleOAuthController(authUsecase usecase.Auth, callbackRedirectURI string) *googleOAuthController {
	return &googleOAuthController{
		authUsecase:         authUsecase,
		callbackRedirectURI: callbackRedirectURI,
	}
}

func (c *googleOAuthController) Callback(w http.ResponseWriter, r *http.Request) {
	slog.Info("googleOAuthController.Callback")

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed url.ParseQuery")
		panic(err)
	}

	code := query.Get("code")
	token, err := c.authUsecase.GetGoogleTokenFromCode(r.Context(), code)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed authUsecase.GetGoogleTokenFromCode", "code", code, "err", err)
		panic(err)
	}

	slog.InfoContext(r.Context(), "token is ", "token", token)

	slog.Info("TODO tokenをcookieに詰め込む")

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, c.callbackRedirectURI, http.StatusFound)
}
