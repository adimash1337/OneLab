package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	Port int `env:"PORT" envDefault:"9000"`
}

func New() (*Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Print("")
	}
	return &config, nil
}
