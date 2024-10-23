package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	Line EnvLineConfig `envPrefix:"LINE_"`
}

type EnvLineConfig struct {
	ChannelSecret string `env:"CHANNEL_SECRET"`
	ChannelToken  string `env:"CHANNEL_TOKEN"`
}

func NewEnvConfig() (*EnvConfig, error) {
	cfg := &EnvConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed env.Parse. err: %w", err)
	}
	return cfg, nil
}
