package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}

func NewService() Service {
	return &svc{}
}
