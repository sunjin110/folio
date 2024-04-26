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
	ctx := r.Context()

	code := query.Get("code")
	output, err := c.authUsecase.StartSessionFromGoogleOAuthCode(ctx, code)
	if err != nil {
		slog.ErrorContext(ctx, "fialed authUsecase.StartSessionFromGoogleOAuthCode", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    output.AccessToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, c.callbackRedirectURI, http.StatusFound)
}
