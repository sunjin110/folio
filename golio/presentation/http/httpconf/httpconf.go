package httpconf

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sunjin110/folio/golio/utils"
)

type Config struct {
	Server         Server            `envPrefix:"SERVER_"`
	GoogleOAuth    GoogleOAuthConfig `envPrefix:"GOOGLE_OAUTH_"`
	SessionKVStore KVStoreConfig     `envPrefix:"SESSION_KV_STORE_"`
	D1Database     D1Config          `envPrefix:"D1_DATABASE_"`
	CORS           CORSConfig        `envPrefix:"CORS_"`
	PostgresDB     PostgresConfig    `envPrefix:"POSTGRES_"`
	AWS            AWSConfig         `envPrefix:"AWS_"`
	MediaS3        S3Config          `envPrefix:"MEDIA_S3_"`
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

type PostgresConfig struct {
	Datasource string `env:"DATASOURCE"`
}

type AWSConfig struct {
	AccessKeyID     string `env:"ACCESS_KEY_ID"`
	SecretAccessKey string `env:"SECRET_ACCESS_KEY"`
	DefaultRegion   string `env:"DEFAULT_REGION"`
}

type S3Config struct {
	Region     string `env:"REGION"`
	BucketName string `env:"BUCKET_NAME"`
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
