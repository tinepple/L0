package orders

import (
	"WBL0/internal/storage"
	"context"
)

type Storage interface {
	GetOrders(ctx context.Context) ([]storage.Order, error)
	GetItems(ctx context.Context, modelId string) ([]storage.Item, error)
}
