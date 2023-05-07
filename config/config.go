package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type (
	Config struct {
		HTTP
		Log
		Postgres
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	Postgres struct {
		MaxOpenConns int    `env:"PG_MAX_OPEN_CONNS"`
		MaxIdleConns int    `env:"PG_MAX_IDLE_CONNS"`
		Host         string `env:"PG_HOST,required"`
		Port         int    `env:"PG_PORT,required"`
		User         string `env:"PG_USER,required"`
		Password     string `env:"PG_PASSWORD,required"`
		DbName       string `env:"PG_DB,required"`
		SSLMode      string `env:"PG_SSL_MODE,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
