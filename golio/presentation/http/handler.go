package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/presentation/http/conv"
	"github.com/sunjin110/folio/golio/usecase"
)

type golioAPIServicer struct {
	articleUsecase usecase.Article
	mediaUsecase   usecase.Media
}

func NewGolioAPIServicer(articleUsecase usecase.Article, mediaUsecase usecase.Media) openapi.GolioAPIServicer {
	return &golioAPIServicer{
		articleUsecase: articleUsecase,
		mediaUsecase:   mediaUsecase,
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

func (g *golioAPIServicer) ArticlesGet(ctx context.Context, offset int32, limit int32) (openapi.ImplResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	output, err := g.articleUsecase.FindSummaries(ctx, offset, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed articleUsecase.FindSummaries", "offset", offset, "limit", limit, "err", err)
	}
	return openapi.Response(http.StatusOK, conv.ToArticlesGet(output.Articles, output.TotalCount)), nil
}

func (g *golioAPIServicer) ArticlesPost(ctx context.Context, req openapi.ArticlesPostRequest) (openapi.ImplResponse, error) {
	if err := g.articleUsecase.Insert(ctx, model.NewArticle(req.Title, req.Body, "", time.Now())); err != nil {
		slog.ErrorContext(ctx, "failed article insert", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, nil), nil
}

func (g *golioAPIServicer) ArticlesArticleIdPut(ctx context.Context, articleID string, req openapi.ArticlesPostRequest) (openapi.ImplResponse, error) {
	if err := g.articleUsecase.Update(ctx, &model.Article{
		ID:        articleID,
		Title:     req.Title,
		Body:      req.Body,
		UpdatedAt: time.Now(),
	}); err != nil {
		slog.ErrorContext(ctx, "failed article update", "err", err, "articleID", articleID, "req", req)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, nil), nil
}

func (g *golioAPIServicer) MediaGet(ctx context.Context, offset int32, limit int32) (openapi.ImplResponse, error) {
	output, err := g.mediaUsecase.FindSummaries(ctx, offset, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed get media", "err", err, "offfset", offset, "limit", limit)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, conv.ToMediaGet(output.Media, output.TotalCount)), nil
}

func (g *golioAPIServicer) MediaMediumIdDelete(ctx context.Context, mediumID string) (openapi.ImplResponse, error) {
	if err := g.mediaUsecase.Delete(ctx, mediumID); err != nil {
		slog.ErrorContext(ctx, "failed delete medium", "err", err, "id", mediumID)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, nil), nil
}
func (g *golioAPIServicer) MediaMediumIdGet(ctx context.Context, mediumID string) (openapi.ImplResponse, error) {
	medium, err := g.mediaUsecase.Get(ctx, mediumID)
	if err != nil {
		slog.ErrorContext(ctx, "fialed get medium", "err", err, "mediumID", mediumID)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}

	if medium == nil {
		return openapi.Response(http.StatusNotFound, nil), nil
	}

	return openapi.Response(http.StatusOK, openapi.MediaMediumIdGet200Response{
		MediumId:     mediumID,
		ThumbnailUrl: medium.ThumbnailURL,
		DownloadUrl:  medium.DownloadURL,
		FileType:     medium.FileType,
	}), nil
}

func (g *golioAPIServicer) MediaPost(ctx context.Context, req openapi.MediaPostRequest) (openapi.ImplResponse, error) {
	presignedURL, err := g.mediaUsecase.Insert(ctx, req.FileName)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}

	return openapi.Response(http.StatusOK, openapi.MediaPost200Response{
		UploadPresignedUrl: presignedURL,
	}), nil
}
