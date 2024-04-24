package http

import (
	"context"
	"net/http"

	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
	"github.com/sunjin110/folio/golio/usecase"
)

func Serve(ctx context.Context, cfg *httpconf.Config) {

	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)
	authUsecase := usecase.NewAuth(googleOAuth2Repo)

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer())

	googleOAuthController := NewGoogleOAuthController(authUsecase, cfg.GoogleOAuth.CallbackRedirectURI)
	r := openapi.NewRouter(golioAPIController)

	// r.HandleFunc("/auth/google-oauth/callback", googleOAuthController.Callback)
	r.Methods(http.MethodGet).Path("/auth/google-oauth/callback").Name("google-oauth/callback").HandlerFunc(googleOAuthController.Callback)

	err := http.ListenAndServe(cfg.Server.PORT, r)
	if err != nil {
		panic(err)
	}
}
