package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppPort   string `env:"APP_PORT" envDefault:":8080"`
	DBHost    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort    string `env:"DB_PORT" envDefault:"5432"`
	DBUser    string `env:"DB_USER" envDefault:"postgres"`
	DBPass    string `env:"DB_PASS" envDefault:"postgres"`
	DBName    string `env:"DB_NAME" envDefault:"spycat"`
	CatAPIURL string `env:"CAT_API_URL" envDefault:"https://api.thecatapi.com/v1"`
	CatAPIKey string `env:"CAT_API_KEY" envDefault:""`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)
}
