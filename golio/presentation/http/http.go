package http

import (
	"net/http"

	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func Serve() {

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer())
	authAPIController := openapi.NewAuthAPIController(nil)
	r := openapi.NewRouter(golioAPIController, authAPIController)
	err := http.ListenAndServe(":8065", r)
	if err != nil {
		panic(err)
	}
}
