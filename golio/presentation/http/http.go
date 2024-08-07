package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/cors"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/aws/dynamodb"
	"github.com/sunjin110/folio/golio/infrastructure/aws/s3"
	"github.com/sunjin110/folio/golio/infrastructure/aws/translate"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt"
	"github.com/sunjin110/folio/golio/infrastructure/gcp/custom_search_api"
	"github.com/sunjin110/folio/golio/infrastructure/postgres"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/dynamodto"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
	"github.com/sunjin110/folio/golio/usecase"
)

func Router(ctx context.Context, cfg *httpconf.Config) (http.Handler, error) {
	googleOAuth2Repo, err := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)
	if err != nil {
		return nil, fmt.Errorf("failed repository.NewGoogleOAuth2. err: %w", err)
	}

	db, err := postgres.OpenDB(cfg.PostgresDB.Datasource)
	if err != nil {
		return nil, fmt.Errorf("failed oepn db: %w", err)
	}

	chatGPTClient := chatgpt.NewClient(cfg.ChatGPT.APIKey)

	googleCustomeSearchClient, err := custom_search_api.NewClient(ctx, cfg.GoogleCustomSearchKey)
	if err != nil {
		return nil, fmt.Errorf("failed googleCustomeSearchClient. err: %w", err)
	}

	googleCustomSearchRepo := repository.NewGoogleCustomSearch(googleCustomeSearchClient)

	htmlContentRepo := repository.NewHtmlContent()

	articleAIRepo := repository.NewArticleAI(chatGPTClient, googleCustomSearchRepo, htmlContentRepo)
	articleTagRepo := repository.NewArticleTag(db)
	articleRepo := repository.NewArticleV2(ctx, db, articleTagRepo)

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed load aws config: %w", err)
	}

	s3Client := s3.NewS3Client(awsCfg)
	if cfg.MediaS3.IsLocakStack {
		s3Client, err = s3.NewLocalStackS3Client()
		if err != nil {
			return nil, fmt.Errorf("failed new local stack s3 client. err: %w", err)
		}
	}

	awsTranslateClient := translate.NewAWSTranslate(awsCfg)

	mediaRepo := repository.NewMedia(db, cfg.MediaS3.BucketName, s3Client)

	dynamoInnerClient := dynamodb.NewInnerClient(awsCfg)
	sessionV3Repo := repository.NewSessionV3(dynamodb.NewClient[dynamodto.UserSessionV3](dynamoInnerClient), cfg.SessionDynamoDB.TableName)
	translateRepo := repository.NewTranslate(awsTranslateClient)
	englishDictionaryRepo := repository.NewEnglishDictionary(cfg.WordsAPI.RapidAPIKey, cfg.WordsAPI.RapidAPIHost)
	userRepo := repository.NewUser(dynamodb.NewClient[dynamodto.User](dynamoInnerClient), cfg.UserDynamoDB.TableName)
	taskRepo := repository.NewTask(db)

	authUsecase := usecase.NewAuth(googleOAuth2Repo, userRepo, sessionV3Repo)
	articleUsecase := usecase.NewArticle(articleRepo, articleAIRepo, articleTagRepo)
	mediaUsecase := usecase.NewMedia(mediaRepo)
	englishDictionaryUsecase := usecase.NewEnglishDictionary(translateRepo, englishDictionaryRepo)
	taskUsecase := usecase.NewTask(taskRepo)

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer(articleUsecase, mediaUsecase, translateRepo, englishDictionaryUsecase, taskUsecase))

	googleOAuthController := NewGoogleOAuthController(authUsecase, cfg.GoogleOAuth.CallbackRedirectURI)
	r := openapi.NewRouter(golioAPIController)

	r.Methods(http.MethodGet).
		Path("/auth/google-oauth").
		Name("google-oauth").
		HandlerFunc(googleOAuthController.Start)

	r.Methods(http.MethodGet).
		Path("/auth/google-oauth/callback").
		Name("google-oauth/callback").
		HandlerFunc(googleOAuthController.Callback)

	r.Methods(http.MethodGet).
		Path("/auth/google-oauth/verify-token-and-start-session").
		Name("google-oauth/verify-token-and-start-session").
		HandlerFunc(googleOAuthController.VerifyTokenAndStartSession)

	// middleware
	r.Use(AuthMW(authUsecase, userRepo))

	// CORSミドルウェアの設定
	// すべてのオリジンからのアクセスを許可する設定
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.GetAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(r), nil
}

func Serve(ctx context.Context, cfg *httpconf.Config) error {
	r, err := Router(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed Router: %w", err)
	}

	slog.Info("server started", "port", cfg.Server.PORT)
	if err := http.ListenAndServe(cfg.Server.PORT, r); err != nil {
		return fmt.Errorf("failed http.ListenAndServe: %w", err)
	}
	return nil
}
