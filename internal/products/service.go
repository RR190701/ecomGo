package products

import (
	"context"

	repo "github.com/RR190701/ecomGo/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	ListOneProduct(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}
func (s *svc) ListOneProduct(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductById(ctx, id)
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}
