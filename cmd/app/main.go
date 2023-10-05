package main

import (
	"context"
	"github.com/Zyigh/hetic-cms/internal/clients"
	"github.com/Zyigh/hetic-cms/internal/config"
	"github.com/Zyigh/hetic-cms/internal/page"
	"github.com/Zyigh/hetic-cms/internal/papi"
	"github.com/Zyigh/hetic-cms/internal/server"
	"log"
)

func run(ctx context.Context) error {
	conf, err := config.New()

	if err != nil {
		return err
	}

	clts, err := clients.New(ctx, conf)

	handlers := []server.Handler{
		page.NewHandler(clts),
		papi.NewHandler(clts),
	}

	httpServer := server.New(conf, handlers)

	if err := httpServer.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
