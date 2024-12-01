package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

type service struct {
	consumer sarama.PartitionConsumer
	storage  Storage
	cache    Cache
}

type Service interface {
	Consume(ctx context.Context) error
}

func New(consumer sarama.PartitionConsumer, storage Storage, cache Cache) Service {
	return &service{
		consumer: consumer,
		storage:  storage,
		cache:    cache,
	}
}
