package company

import (
	"context"
)

type Storage interface {
	Create(context context.Context, company Company) (Company, error)
	Update(ctx context.Context, company Company, id int64) (Company, error)
	Delete(ctx context.Context, id int64) (int64, error)
	List(ctx context.Context) ([]Company, error)
	ListWithFilter(ctx context.Context, f map[string]string) ([]Company, error)
}

//var _ Storage = (*Queries)(nil)
