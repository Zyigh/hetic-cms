package page

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Zyigh/hetic-cms/internal"
	"github.com/Zyigh/hetic-cms/internal/clients"
	"github.com/Zyigh/hetic-cms/internal/entities"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
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

func (r Repository) GetPages(ctx context.Context) ([]entities.Page, error) {
	q, _, err := r.qb.Select("title").From("page").ToSql()

	if err != nil {
		return nil, fmt.Errorf("get pages: %w", err)
	}

	pages := make([]entities.Page, 0)

	if err := pgxscan.Select(ctx, r.client.DB, &pages, q); err != nil {
		return nil, fmt.Errorf("get pages: %w", err)
	}

	return pages, nil
}

func (r Repository) GetOnePage(ctx context.Context, pageID uuid.UUID) (entities.Page, error) {
	q, args, err := r.qb.
		Select("title", "content").
		From("page").
		Where(squirrel.Eq{"id": pageID}).
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
