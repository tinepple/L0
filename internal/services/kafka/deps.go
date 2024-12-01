package kafka

import (
	"WBL0/internal/messages"
	"context"
)

type Storage interface {
	CreateOrder(ctx context.Context, model messages.Order) error
}

type Cache interface {
	Set(key string, value messages.Order) error
}
