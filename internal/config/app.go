package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type App struct {
	Port int    `env:"HTTP_PORT" envDefault:"8080"`
	DB   string `env:"DB_ADDR" envDefault:"postgres://root:root@localhost:5432/page_db?sslmode=disable"`
}

func New() (App, error) {
	app := App{}

	if err := env.Parse(&app); err != nil {
		return app, fmt.Errorf("new app: %w", err)
	}

	return app, nil
}
