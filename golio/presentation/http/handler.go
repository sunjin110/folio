package http

import (
	"context"
	"net/http"

	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

type golioAPIServicer struct {
}

func NewGolioAPIServicer() openapi.GolioAPIServicer {
	return &golioAPIServicer{}
}

func (g *golioAPIServicer) HelloGet(context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, openapi.HelloGet200Response{
		Hello: "world",
	}), nil
}

func (g *golioAPIServicer) ArticlesArticleIdGet(context.Context, string) (openapi.ImplResponse, error) {
	panic("unimplemented")
}

func (g *golioAPIServicer) ArticlesGet(context.Context, string, int32) (openapi.ImplResponse, error) {
	panic("unimplemented")
}

func (g *golioAPIServicer) ArticlesPost(context.Context, openapi.ArticlesPostRequest) (openapi.ImplResponse, error) {
	panic("unimplemented")
}

func (g *golioAPIServicer) ArticlesPut(context.Context, openapi.ArticlesPutRequest) (openapi.ImplResponse, error) {
	panic("unimplemented")
}
