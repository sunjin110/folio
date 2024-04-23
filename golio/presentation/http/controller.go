package http

import (
	"log/slog"
	"net/http"
	"net/url"
)

type googleOAuthController struct {
}

func NewGoogleOAuthController() *googleOAuthController {
	return &googleOAuthController{}
}

func (c *googleOAuthController) Callback(w http.ResponseWriter, r *http.Request) {
	slog.Info("========= callback!!")

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		panic(err)
	}

	query.Get("code")

	// r.ParseForm()

	// TODO なんかあれだよあれしなきゃなんだよ

	http.Redirect(w, r, "http://localhost:3000/login", http.StatusFound)
}
