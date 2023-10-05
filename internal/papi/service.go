package papi

import (
	"context"
	"fmt"
	"github.com/Zyigh/hetic-cms/hetic-cms/facades"
	"github.com/Zyigh/hetic-cms/hetic-cms/models"
	"github.com/Zyigh/hetic-cms/internal/clients"
)

type Service struct {
	repo Repository
}

func NewService(clts clients.Clients) Service {
	return Service{
		repo: NewRepository(clts),
	}
}

func (s Service) GetOnePage(ctx context.Context, facade facades.GetPageForPAPI) (models.PapiPage, error) {
	page, err := s.repo.GetOnePage(ctx, facade.Name)

	if err != nil {
		return models.PapiPage{}, fmt.Errorf("get one page: %w", err)
	}

	return models.PapiPage{
		Name:    page.Title,
		Content: page.Content,
	}, err
}
