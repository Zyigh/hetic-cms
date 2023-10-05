package page

import (
	"context"
	"fmt"
	"github.com/Zyigh/hetic-cms/hetic-cms/facades"
	"github.com/Zyigh/hetic-cms/hetic-cms/models"
	"github.com/Zyigh/hetic-cms/internal"
	"github.com/Zyigh/hetic-cms/internal/clients"
	"github.com/Zyigh/hetic-cms/internal/entities"
)

type Service struct {
	repo Repository
}

func NewService(clts clients.Clients) Service {
	return Service{
		repo: NewRepository(clts),
	}
}

func (s Service) ListPages(ctx context.Context) (models.PagesList, error) {
	pages, err := s.repo.GetPages(ctx)

	if err != nil {
		return nil, fmt.Errorf("list pages: %w", err)
	}

	return internal.Map(pages, func(page entities.Page) models.PageForList {
		return models.PageForList{
			Name: page.Title,
		}
	}), nil
}

func (s Service) GetOnePage(ctx context.Context, facade facades.GetOnePage) (models.SinglePage, error) {
	page, err := s.repo.GetOnePage(ctx, facade.ID)

	if err != nil {
		return models.SinglePage{}, fmt.Errorf("get one page: %w", err)
	}

	return models.SinglePage{
		Name:    page.Title,
		Content: page.Content,
	}, nil
}
