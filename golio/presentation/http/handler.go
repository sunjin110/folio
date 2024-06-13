package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/presentation/http/conv"
	"github.com/sunjin110/folio/golio/usecase"
)

type golioAPIServicer struct {
	articleUsecase usecase.Article
	mediaUsecase   usecase.Media
	translateRepo  repository.Translate
}

func NewGolioAPIServicer(articleUsecase usecase.Article, mediaUsecase usecase.Media, translateRepo repository.Translate) openapi.GolioAPIServicer {
	return &golioAPIServicer{
		articleUsecase: articleUsecase,
		mediaUsecase:   mediaUsecase,
		translateRepo:  translateRepo,
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
		Tags:      conv.ToArticleTags(article.Tags),
	}), nil
}

func (g *golioAPIServicer) ArticlesGet(ctx context.Context, offset int32, limit int32, titleSearchTextQueryParam string) (openapi.ImplResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	var titleSearchText *string
	if titleSearchTextQueryParam != "" {
		titleSearchText = &titleSearchTextQueryParam
	}

	output, err := g.articleUsecase.FindSummaries(ctx, offset, limit, titleSearchText)
	if err != nil {
		slog.ErrorContext(ctx, "failed articleUsecase.FindSummaries", "offset", offset, "limit", limit, "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, conv.ToArticlesGet(output.Articles, output.TotalCount)), nil
}

func (g *golioAPIServicer) ArticlesPost(ctx context.Context, req openapi.ArticlesPostRequest) (openapi.ImplResponse, error) {
	inserted, err := g.articleUsecase.Insert(ctx, model.NewArticle(req.Title, req.Body, "", time.Now()))
	if err != nil {
		slog.ErrorContext(ctx, "failed article insert", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, openapi.ArticlesPost200Response{
		Id: inserted.ID,
	}), nil
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

func (g *golioAPIServicer) ArticlesArticleIdAiPut(ctx context.Context, articleID string, req openapi.ArticlesArticleIdAiPutRequest) (openapi.ImplResponse, error) {
	article, err := g.articleUsecase.AssistantBodyByAI(ctx, articleID, req.Message)
	if err != nil {
		slog.ErrorContext(ctx, "failed articleUsecase.AssistantBodyByAI", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal server error"), nil
	}
	if article == nil {
		return openapi.Response(http.StatusNotFound, "not found"), nil
	}

	return openapi.Response(http.StatusOK, openapi.ArticlesArticleIdAiPut200Response{
		GeneratedBody: article.Body,
	}), nil
}

func (g *golioAPIServicer) ArticlesAiPost(ctx context.Context, req openapi.ArticlesAiPostRequest) (openapi.ImplResponse, error) {
	article, err := g.articleUsecase.GenerateArticleByAI(ctx, req.Prompt)
	if err != nil {
		slog.ErrorContext(ctx, "failed articleUsecase.GenerateArticleByAI", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal server error"), nil
	}

	return openapi.Response(http.StatusOK, openapi.ArticlesAiPost200Response{
		ArticleId: article.ID,
	}), nil
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

func (g *golioAPIServicer) TranslationPost(ctx context.Context, req openapi.TranslationPostRequest) (openapi.ImplResponse, error) {
	translatedText, err := g.translateRepo.TranslateText(ctx, req.Text, model.TranslateLanguageCode(req.SourceLanguageCode), model.TranslateLanguageCode(req.TargetLanguageCode))
	if err != nil {
		slog.ErrorContext(ctx, "failed translate text", "req", req, "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, openapi.TranslationPost200Response{
		TranslatedText: translatedText,
	}), nil
}

func (g *golioAPIServicer) ArticlesTagsGet(ctx context.Context, searchText string, offset int32, limit int32) (openapi.ImplResponse, error) {
	var nameSearchText *string
	if searchText != "" {
		nameSearchText = &searchText
	}

	tags, err := g.articleUsecase.FindTags(ctx, offset, limit, nameSearchText)
	if err != nil {
		slog.ErrorContext(ctx, "failed find tags", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}

	return openapi.Response(http.StatusOK, openapi.ArticlesTagsGet200Response{
		Tags: conv.ToArticleTags(tags),
	}), nil
}

func (g *golioAPIServicer) ArticlesTagsPost(ctx context.Context, req openapi.ArticlesTagsPostRequest) (openapi.ImplResponse, error) {
	articleTag := model.NewArticleTag(req.Name, time.Now())
	if err := g.articleUsecase.InsertTag(ctx, articleTag); err != nil {
		slog.ErrorContext(ctx, "failed insert tag", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, openapi.InsertArticleTagResponse{
		Id: articleTag.ID,
	}), nil
}

func (g *golioAPIServicer) ArticlesTagsTagIdDelete(ctx context.Context, tagID string) (openapi.ImplResponse, error) {
	if err := g.articleUsecase.DeleteTag(ctx, tagID); err != nil {
		slog.ErrorContext(ctx, "failed delete tag", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, openapi.DeleteArticleTagResponse{
		Id: tagID,
	}), nil
}

func (g *golioAPIServicer) ArticlesTagsTagIdPut(ctx context.Context, tagID string, req openapi.ArticlesTagsTagIdPutRequest) (openapi.ImplResponse, error) {
	if err := g.articleUsecase.UpdateTag(ctx, &model.ArticleTag{
		ID:          tagID,
		Name:        req.Name,
		UpdatedTime: time.Now(),
	}); err != nil {
		slog.ErrorContext(ctx, "failed article tag update", "err", err)
		return openapi.Response(http.StatusInternalServerError, "internal"), nil
	}
	return openapi.Response(http.StatusOK, openapi.UpdateArticleTagResponse{
		Id: tagID,
	}), nil
}
