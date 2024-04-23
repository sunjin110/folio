package http

import (
	"context"
	"net/http"

	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/usecase"
)

func Serve(ctx context.Context) {

	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, "client_id", "client_secret", "redirect_url")
	authUsecase := usecase.NewAuth(googleOAuth2Repo)

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer())
	authAPIController := openapi.NewAuthAPIController(nil)

	googleOAuthController := NewGoogleOAuthController(authUsecase)
	r := openapi.NewRouter(golioAPIController, authAPIController)

	r.HandleFunc("/auth/google-oauth/callback", googleOAuthController.Callback)

	err := http.ListenAndServe(":3001", r)
	if err != nil {
		panic(err)
	}
}
