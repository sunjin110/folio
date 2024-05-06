package lambdaconf

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sunjin110/folio/golio/utils"
)

type Config struct {
	GoogleOAuth    GoogleOAuthConfig `envPrefix:"GOOGLE_OAUTH_"`
	SessionKVStore KVStoreConfig     `envPrefix:"SESSION_KV_STORE_"`
	D1Database     D1Config          `envPrefix:"D1_DATABASE_"`
	CORS           CORSConfig        `envPrefix:"CORS_"`
}

type GoogleOAuthConfig struct {
	ClientID            string `env:"CLIENT_ID"`
	ClientSecret        string `env:"CLIENT_SECRET"`
	RedirectURI         string `env:"REDIRECT_URI"`
	CallbackRedirectURI string `env:"CALLBACK_REDIRECT_URI"`
}

type KVStoreConfig struct {
	AccountID   string `env:"ACCOUNT_ID"`
	NamespaceID string `env:"NAMESPACE_ID"`
	APIToken    string `env:"API_TOKEN"`
}

type D1Config struct {
	AccountID  string `env:"ACCOUNT_ID"`
	DatabaseID string `env:"DATABASE_ID"`
	APIToken   string `env:"API_TOKEN"`
}

type CORSConfig struct {
	AllowedOrigins string `env:"ALLOWED_ORIGINS"`
}

func (c *CORSConfig) GetAllowedOrigins() []string {
	return strings.Split(c.AllowedOrigins, ",")
}

func NewConfig() (*Config, error) {
	envFilePath := filepath.Join(utils.ProjectRoot(), ".env")
	_ = godotenv.Load(envFilePath)
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse basic config: %w", err)
	}
	return cfg, nil
}
