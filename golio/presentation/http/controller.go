package http

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/sunjin110/folio/golio/usecase"
)

type googleOAuthController struct {
	authUsecase usecase.Auth
}

func NewGoogleOAuthController(authUsecase usecase.Auth) *googleOAuthController {
	return &googleOAuthController{
		authUsecase: authUsecase,
	}
}

func (c *googleOAuthController) Callback(w http.ResponseWriter, r *http.Request) {
	slog.Info("========= callback!!")

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed url.ParseQuery")
		panic(err)
	}

	code := query.Get("code")
	token, err := c.authUsecase.GetGoogleTokenFromCode(r.Context(), code)
	if err != nil {
		panic(err)
	}

	slog.InfoContext(r.Context(), "token is ", "token", token)

	slog.Info("TODO tokenをcookieに詰め込む")

	http.Redirect(w, r, "http://localhost:3000/login", http.StatusFound)
}