package lambda

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/rs/cors"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/aws/dynamodb"
	"github.com/sunjin110/folio/golio/infrastructure/aws/s3"
	"github.com/sunjin110/folio/golio/infrastructure/aws/translate"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt"
	"github.com/sunjin110/folio/golio/infrastructure/postgres"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/dynamodto"
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
	if err := postgres.MigrateDB(cfg.PostgresDB.Datasource); err != nil {
		return fmt.Errorf("failed postgres migrate: %w", err)
	}
	return nil
}

func GetHandler(ctx context.Context) (lambdaHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), err error) {
	cfg := lambdaConfig
	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)

	db, err := postgres.OpenDB(cfg.PostgresDB.Datasource)
	if err != nil {
		return nil, fmt.Errorf("failed open db: %w", err)
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed load aws config: %w", err)
	}

	awsTranslateClient := translate.NewAWSTranslate(awsCfg)

	chatGPTClient := chatgpt.NewClient(cfg.ChatGPT.APIKey)

	articleRepo := repository.NewArticleV2(ctx, db, chatGPTClient)

	mediaRepo := repository.NewMedia(db, cfg.MediaS3.BucketName, s3.NewS3Client(awsCfg))

	dynamoInnerClient := dynamodb.NewInnerClient(awsCfg)
	sessionV2Repo := repository.NewSessionV2(dynamodb.NewClient[dynamodto.UserSessionV2](dynamoInnerClient), cfg.SessionDynamoDB.TableName)

	authUsecase := usecase.NewAuth(googleOAuth2Repo, sessionV2Repo)
	articleUsecase := usecase.NewArticle(articleRepo)
	mediaUsecase := usecase.NewMedia(mediaRepo)

	translateRepo := repository.NewTranslate(awsTranslateClient)

	golioAPIController := openapi.NewGolioAPIController(golio_http.NewGolioAPIServicer(articleUsecase, mediaUsecase, translateRepo))

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
	r.Use(golio_http.AuthMW(authUsecase))

	// CORSミドルウェアの設定
	// すべてのオリジンからのアクセスを許可する設定
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.GetAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Cookie"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	adapter := httpadapter.New(handler)

	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		reqEndLogFunc := accessLog(ctx, req)
		defer reqEndLogFunc()

		return adapter.ProxyWithContext(ctx, req)
	}, nil
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
