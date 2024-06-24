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
	GoogleOAuth           GoogleOAuthConfig `envPrefix:"GOOGLE_OAUTH_"`
	SessionDynamoDB       DynamoDBConfig    `envPrefix:"SESSION_DYNAMODB_"`
	CORS                  CORSConfig        `envPrefix:"CORS_"`
	PostgresDB            PostgresConfig    `envPrefix:"POSTGRES_"`
	AWS                   AWSConfig         `envPrefix:"AWS_"`
	MediaS3               S3Config          `envPrefix:"MEDIA_S3_"`
	ChatGPT               ChatGPTConfig     `envPrefix:"CHAT_GPT_"`
	GoogleCustomSearchKey string            `env:"GOOGLE_CUSTOM_SEARCH_KEY"`
	WordsAPI              WordsAPIConfig    `envPrefix:"WORDS_API_"`
}

type GoogleOAuthConfig struct {
	ClientID            string `env:"CLIENT_ID"`
	ClientSecret        string `env:"CLIENT_SECRET"`
	RedirectURI         string `env:"REDIRECT_URI"`
	CallbackRedirectURI string `env:"CALLBACK_REDIRECT_URI"`
}

type CORSConfig struct {
	AllowedOrigins string `env:"ALLOWED_ORIGINS"`
}

type DynamoDBConfig struct {
	TableName string `env:"TABLE_NAME"`
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

type ChatGPTConfig struct {
	APIKey string `env:"API_KEY"`
}

type WordsAPIConfig struct {
	RapidAPIKey  string `env:"RAPID_API_KEY"`
	RapidAPIHost string `env:"RAPID_API_HOST"`
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
