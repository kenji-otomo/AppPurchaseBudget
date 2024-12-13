package config

import (
	"github.com/caarlos0/env"
)

type config struct {
	Dns string `env:"DNS"`
}

var cfg config

func Load() (*config, error) {
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
