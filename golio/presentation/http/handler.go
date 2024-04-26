package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/usecase"
)

type golioAPIServicer struct {
	articleUsecase usecase.Article
}

func NewGolioAPIServicer(articleUsecase usecase.Article) openapi.GolioAPIServicer {
	return &golioAPIServicer{
		articleUsecase: articleUsecase,
	}
}

func (g *golioAPIServicer) HelloGet(context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, openapi.HelloGet200Response{
		Hello: "world",
	}), nil
}

func (g *golioAPIServicer) ArticlesArticleIdGet(ctx context.Context, articleID string) (openapi.ImplResponse, error) {
	article, err := g.articleUsecase.Get(ctx, articleID)
	if err != nil {
		slog.ErrorContext(ctx, "failed articleUsecase.Get", articleID, err)
		return openapi.Response(http.StatusInternalServerError, "internal server error"), nil
	}
	if article == nil {
		return openapi.Response(http.StatusNotFound, "not found"), nil
	}

	return openapi.Response(http.StatusOK, openapi.ArticlesArticleIdGet200Response{
		Id:        articleID,
		Title:     article.Title,
		Body:      article.Body,
		CreatedAt: article.CreatedAt,
		UserId:    "todo",
	}), nil
}

func (g *golioAPIServicer) ArticlesGet(context.Context, string, int32) (openapi.ImplResponse, error) {
	panic("unimplemented")
}

func (g *golioAPIServicer) ArticlesPost(ctx context.Context, req openapi.ArticlesPostRequest) (openapi.ImplResponse, error) {
	if err := g.articleUsecase.Insert(ctx, model.NewArticle(req.Title, req.Body, "", time.Now())); err != nil {
		slog.ErrorContext(ctx, "failed article insert", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, nil), nil
}

func (g *golioAPIServicer) ArticlesPut(ctx context.Context, req openapi.ArticlesPutRequest) (openapi.ImplResponse, error) {
	if err := g.articleUsecase.Update(ctx, &model.Article{
		ID:        req.Id,
		Title:     req.Title,
		Body:      req.Body,
		UpdatedAt: time.Now(),
	}); err != nil {
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, nil), nil
}
