package httpconf

import (
	"fmt"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sunjin110/folio/golio/utils"
)

type Config struct {
	Server                Server            `envPrefix:"SERVER_"`
	GoogleOAuth           GoogleOAuthConfig `envPrefix:"GOOGLE_OAUTH_"`
	AuthenticationKVStore KVStoreConfig     `envPrefix:"AUTHENTICATION_KV_STORE_"`
}

type Server struct {
	PORT string `env:"PORT"`
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

func NewConfig() (*Config, error) {
	envFilePath := filepath.Join(utils.ProjectRoot(), ".env")
	_ = godotenv.Load(envFilePath)
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("fialed to parse basic config: %w", err)
	}
	return cfg, nil
}