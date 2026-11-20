package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/RR190701/ecomGo/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	CreateOrder(ctx context.Context, order createOrderParams) (repo.Order, error)
}

type svc struct {
	repo repo.Querier
	db   *pgx.Conn
}

var (
	ErrProductNotFound = errors.New("Product not found")

	ErrProductHasNoStock = errors.New("Product has no stock")
)

func (s *svc) CreateOrder(ctx context.Context, order createOrderParams) (repo.Order, error) {

	//basic validations

	if order.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("Customer ID is required")
	}
	if Items := order.Items; len(Items) == 0 {
		return repo.Order{}, fmt.Errorf("Order must have at least one item")
	}
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, fmt.Errorf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)
	//WithTx is not part of the Querier interface, so build tx-scoped queries directly
	qtx := repo.New(tx)

	createdOrder, err := qtx.CreateOrder(ctx, order.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range order.Items {
		product, err := qtx.FindProductById(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductHasNoStock

		}
		// TODO: insert the order item once a CreateOrderItem query exists in queries.sql
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    createdOrder.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCenters,
		})

		if err != nil {
			return repo.Order{}, err
		}

	}

	if err := tx.Commit(ctx); err != nil {
		return repo.Order{}, fmt.Errorf("Failed to commit transaction: %v", err)
	}

	return createdOrder, nil
}

func NewService(repo repo.Querier, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}
