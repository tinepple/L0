package kafka

import (
	"WBL0/internal/messages"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *service) Consume(ctx context.Context) error {
	for {
		select {
		case msg, ok := <-s.consumer.Messages():
			if !ok {
				return errors.New("channel closed, exiting")
			}

			var order messages.Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				fmt.Println("Error unmarshaling JSON", err.Error())
				continue
			}

			err = s.storage.CreateOrder(ctx, order)
			if err != nil {
				fmt.Println("Error creating order", err.Error())
				continue
			}

			s.cache.Set(order.OrderUID, order)
		}
	}
}
