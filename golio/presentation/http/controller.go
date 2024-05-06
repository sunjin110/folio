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

func (c *googleOAuthController) Start(w http.ResponseWriter, r *http.Request) {
	slog.Info("googleOAuthController.Start")
	url, err := c.authUsecase.GenerateGoogleAuthorizationURL()
	if err != nil {
		slog.ErrorContext(r.Context(), "failed authUsecase.GenerateGoogleAuthorizationURL", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func (c *googleOAuthController) Callback(w http.ResponseWriter, r *http.Request) {
	slog.Info("googleOAuthController.Callback")
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed url.ParseQuery")
		http.Error(w, "invalid query", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	code := query.Get("code")
	if code == "" {
		slog.Info("code is empty")
		http.Error(w, "invalid query: code is required", http.StatusBadRequest)
		return
	}

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
		Secure:   true,
		Path:     "/",
		// SameSite: http.SameSiteLaxMode,
		SameSite: http.SameSiteNoneMode,
	})
	http.Redirect(w, r, c.callbackRedirectURI, http.StatusFound)
}
