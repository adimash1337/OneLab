package config

import (
	"awesomeProject/internal/logger"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port string `env:"PORT" envDefault:":9000"`
	DB   string `env:"DB" envDefault:"in-memory"`
}

func New() *Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		logger.Logger().Println(err)

	}
	return &config
}
