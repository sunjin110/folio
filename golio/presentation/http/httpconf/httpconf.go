package httpconf

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sunjin110/folio/golio/utils"
)

type Config struct {
	PORT        string            `env:"PORT"`
	GoogleOAuth GoogleOAuthConfig `envPrefix:"GOOGLE_OAUTH_"`
}

type Server struct {
}

type GoogleOAuthConfig struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectURI  string `env:"REDIRECT_URI"`
}

func NewConfig() (*Config, error) {

	envFilePath := filepath.Join(utils.ProjectRoot(), ".env")
	slog.Info("a", "envFilePath", envFilePath)

	err := godotenv.Load(envFilePath)
	slog.Error("godotenv.Load error", "err", err)

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("fialed to parse basic config: %w", err)
	}
	return cfg, nil
}
