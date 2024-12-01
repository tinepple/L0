package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

type service struct {
	consumer sarama.PartitionConsumer
	storage  Storage
	handler  Handler
}

type Service interface {
	Consume(ctx context.Context) error
}

func New(consumer sarama.PartitionConsumer, storage Storage, handler Handler) Service {
	return &service{
		consumer: consumer,
		storage:  storage,
		handler:  handler,
	}
}
