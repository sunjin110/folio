package lambda

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	golio_http "github.com/sunjin110/folio/golio/presentation/http"
	"github.com/sunjin110/folio/golio/presentation/lambda/lambdaconf"
	"github.com/sunjin110/folio/golio/usecase"
)

var lambdaConfig *lambdaconf.Config

func Setup() error {
	cfg, err := lambdaconf.NewConfig()
	if err != nil {
		return fmt.Errorf("failed lambdaconf.NewConfig: %w", err)
	}
	lambdaConfig = cfg
	return nil
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	reqEndLogFunc := accessLog(ctx, request)
	defer reqEndLogFunc()

	cfg := lambdaConfig

	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)
	sessionRepo, err := repository.NewSessionKVStore(ctx, cfg.SessionKVStore.APIToken, cfg.SessionKVStore.AccountID, cfg.SessionKVStore.NamespaceID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed repository.NewSessionKVStore: %w", err)
	}

	d1Client, err := d1.NewClient(cfg.D1Database.AccountID, cfg.D1Database.DatabaseID, cfg.D1Database.APIToken)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed d1.NewClient: %w", err)
	}

	articleRepo, err := repository.NewArticle(ctx, d1Client)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed repository.NewArticle: %w", err)
	}

	authUsecase := usecase.NewAuth(googleOAuth2Repo, sessionRepo)
	articleUsecase := usecase.NewArticle(articleRepo)

	golioAPIController := openapi.NewGolioAPIController(golio_http.NewGolioAPIServicer(articleUsecase))

	googleOAuthController := golio_http.NewGoogleOAuthController(authUsecase, cfg.GoogleOAuth.CallbackRedirectURI)
	r := openapi.NewRouter(golioAPIController)

	r.Methods(http.MethodGet).
		Path("/auth/google-oauth").
		Name("google-oauth").
		HandlerFunc(googleOAuthController.Start)

	r.Methods(http.MethodGet).
		Path("/auth/google-oauth/callback").
		Name("google-oauth/callback").
		HandlerFunc(googleOAuthController.Callback)

	// middleware
	r.Use(golio_http.CorsMW())
	r.Use(golio_http.AuthMW(authUsecase))

	adapter := gorillamux.New(r)

	res, err := adapter.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&request))
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed adapter.ProxyWithContext: %w", err)
	}

	return *res.Version1(), nil
}

func accessLog(ctx context.Context, request events.APIGatewayProxyRequest) (reqEndLogFunc func()) {

	// access log
	startTime := time.Now()

	slog.InfoContext(ctx, "request start",
		"start_time", startTime,
		"method", request.HTTPMethod,
		"path", request.Path)

	return func() {
		slog.Info("request end", "end_time", time.Now(), "req_duration_milli", time.Since(startTime).Milliseconds())
	}
}
