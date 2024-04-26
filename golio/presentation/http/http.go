package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rs/cors"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
	"github.com/sunjin110/folio/golio/usecase"
)

func Serve(ctx context.Context, cfg *httpconf.Config) error {
	googleOAuth2Repo := repository.NewGoogleOAuth2(ctx, cfg.GoogleOAuth.ClientID, cfg.GoogleOAuth.ClientSecret, cfg.GoogleOAuth.RedirectURI)
	sessionRepo, err := repository.NewSessionKVStore(ctx, cfg.SessionKVStore.APIToken, cfg.SessionKVStore.AccountID, cfg.SessionKVStore.NamespaceID)
	if err != nil {
		return fmt.Errorf("failed repository.NewSessionKVStore: %w", err)
	}

	d1Client, err := d1.NewClient(cfg.D1Database.AccountID, cfg.D1Database.DatabaseID, cfg.D1Database.APIToken)
	if err != nil {
		return fmt.Errorf("failed d1.NewClient: %w", err)
	}

	articleRepo, err := repository.NewArticle(ctx, d1Client)
	if err != nil {
		return fmt.Errorf("failed repository.NewArticle: %w", err)
	}

	authUsecase := usecase.NewAuth(googleOAuth2Repo, sessionRepo)
	articleUsecase := usecase.NewArticle(articleRepo)

	golioAPIController := openapi.NewGolioAPIController(NewGolioAPIServicer(articleUsecase))

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

	// CORSミドルウェアの設定
	// すべてのオリジンからのアクセスを許可する設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // すべてのオリジンを許可 TODO 限られたoriginのみにあとで変更する
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	handler := c.Handler(r)

	slog.Info("server started", "port", cfg.Server.PORT)
	if err := http.ListenAndServe(cfg.Server.PORT, handler); err != nil {
		return fmt.Errorf("failed http.ListenAndServe: %w", err)
	}
	return nil
}
