package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App     `yaml:"app"`
		HTTP    `yaml:"http"`
		Log     `yaml:"logger"`
		PG      `yaml:"postgres"`
		GRPC    `yaml:"grpc"`
		Rabbit  `yaml:"rabbit"`
		Graylog `yaml:"graylog"`
		Swagger `yaml:"swagger"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		SessionHttpKey string `env-required:"true" yaml:"session_http_key" env:"SESSION_HTTP_KEY"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	GRPC struct {
		URL string `env-required:"true" yaml:"url" env:"GRPC_URL"`
	}

	Graylog struct {
		URL string `env-required:"true" yaml:"url" env:"GRAYLOG_URL"`
	}

	Swagger struct {
		URL string `env-required:"true" yaml:"url" env:"SWAGGER_URL"`
	}

	// Rabbit -.
	Rabbit struct {
		URL      string `env-required:"true" yaml:"url"       env:"RABBIT_URL"`
		ClientID string `env-required:"true" yaml:"client_id" env:"RABBIT_CLIENT_ID"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		//return nil, fmt.Errorf("config error: %w", err)
	}
	// cfg.PG.URL = "postgres://user:pass@localhost:5432/postgres"
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
