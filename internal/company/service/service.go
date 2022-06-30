package service

import (
	"context"
	"github.com/yerlanov/xmexercise/internal/company"
)

type service struct {
	storage company.Storage
}

func NewService(storage company.Storage) Service {
	return &service{storage: storage}
}

type Service interface {
	Create(ctx context.Context, dto company.Company) (company.Company, error)
	Delete(ctx context.Context, id int64) (int64, error)
	Update(ctx context.Context, company company.Company, id int64) (company.Company, error)
	List(ctx context.Context) ([]company.Company, error)
	ListWithFilter(ctx context.Context, filter map[string]string) ([]company.Company, error)
}

func (s *service) Create(ctx context.Context, company company.Company) (company.Company, error) {
	return s.storage.Create(ctx, company)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	return s.storage.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, company company.Company, id int64) (company.Company, error) {
	return s.storage.Update(ctx, company, id)
}

func (s *service) List(ctx context.Context) ([]company.Company, error) {
	return s.storage.List(ctx)
}

func (s *service) ListWithFilter(ctx context.Context, filter map[string]string) ([]company.Company, error) {
	withFilter, err := s.storage.ListWithFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return withFilter, nil
}
