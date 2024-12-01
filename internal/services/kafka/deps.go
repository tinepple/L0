package kafka

import (
	"WBL0/internal/services/orders"
	"WBL0/internal/storage"
	"context"
)

type Storage interface {
	GetOrders(ctx context.Context) ([]storage.Order, error)
	GetItems(ctx context.Context, modelId string) ([]storage.Item, error)
	CreateOrder(ctx context.Context, model storage.Message) error
}

type Handler interface {
	AddModelToCache(order orders.OrderResponse)
}
