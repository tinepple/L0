package kafka

import (
	"WBL0/internal/services/orders"
	"WBL0/internal/storage"
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

			fmt.Println("получили сообщение")

			var receivedMessage Message
			err := json.Unmarshal(msg.Value, &receivedMessage)
			if err != nil {
				fmt.Println("Error unmarshaling JSON", err.Error())
				continue
			}

			fmt.Println("выводим сообщение", receivedMessage)

			var items []storage.Item

			for _, item := range receivedMessage.Items {
				items = append(items, storage.Item(item))
			}

			err = s.storage.CreateOrder(ctx, storage.Message{
				OrderUID:          receivedMessage.OrderUID,
				TrackNumber:       receivedMessage.TrackNumber,
				Entry:             receivedMessage.Entry,
				Delivery:          storage.Delivery(receivedMessage.Delivery),
				Payment:           storage.Payment(receivedMessage.Payment),
				Items:             items,
				Locale:            receivedMessage.Locale,
				InternalSignature: receivedMessage.InternalSignature,
				CustomerID:        receivedMessage.CustomerID,
				DeliveryService:   receivedMessage.DeliveryService,
				Shardkey:          receivedMessage.Shardkey,
				SmID:              receivedMessage.SmID,
				DateCreated:       receivedMessage.DateCreated,
				OofShard:          receivedMessage.OofShard,
			})
			if err != nil {
				fmt.Println("Error creating order", err.Error())
				continue
			}

			var itemsOrder []orders.Item

			for _, item := range receivedMessage.Items {
				itemsOrder = append(itemsOrder, orders.Item(item))
			}

			s.handler.AddModelToCache(orders.OrderResponse{
				OrderUID:          receivedMessage.OrderUID,
				TrackNumber:       receivedMessage.TrackNumber,
				Entry:             receivedMessage.Entry,
				Delivery:          orders.Delivery(receivedMessage.Delivery),
				Payment:           orders.Payment(receivedMessage.Payment),
				Items:             itemsOrder,
				Locale:            receivedMessage.Locale,
				InternalSignature: receivedMessage.InternalSignature,
				CustomerID:        receivedMessage.CustomerID,
				DeliveryService:   receivedMessage.DeliveryService,
				Shardkey:          receivedMessage.Shardkey,
				SmID:              receivedMessage.SmID,
				DateCreated:       receivedMessage.DateCreated,
				OofShard:          receivedMessage.OofShard,
			})
		}
	}
}
