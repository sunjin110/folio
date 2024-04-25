package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
	"github.com/sunjin110/folio/golio/usecase"
)

func Serve(ctx context.Context, cfg *httpconf.Config) error {
	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)
	authRepo, err := repository.NewAuthorizationKVStore(ctx, "api_token", "account_id", "namespace")
	if err != nil {
		return fmt.Errorf("failed repository.NewAuthorizationKVStore: %w", err)
	}

	authUsecase := usecase.NewAuth(googleOAuth2Repo, authRepo)

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer())

	googleOAuthController := NewGoogleOAuthController(authUsecase, cfg.GoogleOAuth.CallbackRedirectURI)
	r := openapi.NewRouter(golioAPIController)

	r.Methods(http.MethodGet).Path("/auth/google-oauth/callback").Name("google-oauth/callback").HandlerFunc(googleOAuthController.Callback)

	if err := http.ListenAndServe(cfg.Server.PORT, r); err != nil {
		return fmt.Errorf("failed http.ListenAndServe: %w", err)
	}
	return nil
}
