package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/cors"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/aws/s3"
	"github.com/sunjin110/folio/golio/infrastructure/postgres"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
	"github.com/sunjin110/folio/golio/usecase"
)

func Router(ctx context.Context, cfg *httpconf.Config) (http.Handler, error) {
	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)
	sessionRepo, err := repository.NewSessionKVStore(ctx, cfg.SessionKVStore.APIToken, cfg.SessionKVStore.AccountID, cfg.SessionKVStore.NamespaceID)
	if err != nil {
		return nil, fmt.Errorf("failed repository.NewSessionKVStore: %w", err)
	}

	// d1Client, err := d1.NewClient(cfg.D1Database.AccountID, cfg.D1Database.DatabaseID, cfg.D1Database.APIToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed d1.NewClient: %w", err)
	// }

	db, err := postgres.OpenDB(cfg.PostgresDB.Datasource)
	if err != nil {
		return nil, fmt.Errorf("failed oepn db: %w", err)
	}

	// articleRepo, err := repository.NewArticle(ctx, d1Client)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed repository.NewArticle: %w", err)
	// }

	articleRepo := repository.NewArticleV2(ctx, db)

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed load aws config: %w", err)
	}

	mediaRepo := repository.NewMedia(db, cfg.MediaS3.BucketName, s3.NewS3Client(awsCfg))

	authUsecase := usecase.NewAuth(googleOAuth2Repo, sessionRepo)
	articleUsecase := usecase.NewArticle(articleRepo)
	mediaUsecase := usecase.NewMedia(mediaRepo)

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer(articleUsecase, mediaUsecase))

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

	// middleware
	r.Use(AuthMW(authUsecase))

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
