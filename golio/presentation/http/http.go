package http

import (
	"net/http"

	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func Serve() {

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer())
	authAPIController := openapi.NewAuthAPIController(nil)

	googleOAuthController := NewGoogleOAuthController()
	r := openapi.NewRouter(golioAPIController, authAPIController)

	r.HandleFunc("/auth/google-oauth/callback", googleOAuthController.Callback)

	err := http.ListenAndServe(":3001", r)
	if err != nil {
		panic(err)
	}
}
