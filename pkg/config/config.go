package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// server init config
type ServerConfig struct {
	Port int `env:"PORT" envDefault:"3000"`
}

type PGConfig struct {
	Host     string `env:"PG_HOST" envDefault:"localhost"`
	Port     string `env:"PG_PORT" envDefault:"5434"`
	User     string `env:"PG_USER" envDefault:"postgres"`
	DBNAME   string `env:"PG_DB_NAME" envDefault:"todo_db"`
	Password string `env:"PG_PASS" envDefault:"postgres"`
}

type Config struct {
	ServerConfig ServerConfig
	PGConfig     PGConfig
}

// load config
func Load() (*Config, error) {
	// make a pointer of Config
	cfg := &Config{}
	// if the default values are not set, load from .env
	// actually that ensures critical values are set
	opt := env.Options{
		RequiredIfNoDef: true,
	}

	// parse values into config
	if err := env.ParseWithOptions(cfg, opt); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
