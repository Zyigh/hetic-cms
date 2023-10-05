package clients

import (
	"context"
	"fmt"
	"github.com/Zyigh/hetic-cms/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Clients struct {
	DB *pgxpool.Pool
}

func New(ctx context.Context, conf config.App) (Clients, error) {
	pool, err := pgxpool.New(ctx, conf.DB)

	if err != nil {
		return Clients{}, fmt.Errorf("new clients: %w", err)
	}

	return Clients{DB: pool}, nil
}
