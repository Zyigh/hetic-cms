package papi

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Zyigh/hetic-cms/internal"
	"github.com/Zyigh/hetic-cms/internal/clients"
	"github.com/Zyigh/hetic-cms/internal/entities"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	client clients.Clients
	qb     squirrel.StatementBuilderType
}

func NewRepository(clts clients.Clients) Repository {
	return Repository{
		client: clts,
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r Repository) GetOnePage(ctx context.Context, pageName string) (entities.Page, error) {
	q, args, err := r.qb.
		Select("title", "content").
		From("page").
		Where(squirrel.Eq{"title": pageName}).
		ToSql()

	if err != nil {
		return entities.Page{}, fmt.Errorf("get one page: %w", err)
	}

	page := entities.Page{}

	if err := pgxscan.Get(ctx, r.client.DB, &page, q, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Page{}, internal.ErrNoRows
		}
		return entities.Page{}, fmt.Errorf("get one page: %w", err)
	}

	return page, nil
}
