package orders

import "context"

func (s *Service) GetOrders(ctx context.Context) ([]OrderResponse, error) {
	orders, err := s.storage.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	var result []OrderResponse
	for _, order := range orders {
		items, err := s.storage.GetItems(ctx, order.OrderUID)
		if err != nil {
			return nil, err
		}

		resultItems := make([]Item, 0, len(items))
		for _, item := range items {
			resultItems = append(resultItems, Item{
				ChrtID:      item.ChrtID,
				TrackNumber: item.TrackNumber,
				Price:       item.Price,
				Rid:         item.Rid,
				Name:        item.Name,
				Sale:        item.Sale,
				Size:        item.Size,
				TotalPrice:  item.TotalPrice,
				NmID:        item.Status,
				Brand:       item.Brand,
				Status:      item.Status,
			})
		}

		result = append(result, OrderResponse{
			OrderUID:    order.OrderUID,
			TrackNumber: order.TrackNumber,
			Entry:       order.Entry,
			Delivery: Delivery{
				Name:    order.Name,
				Phone:   order.Phone,
				Zip:     order.Zip,
				City:    order.City,
				Address: order.Address,
				Region:  order.Region,
				Email:   order.Email,
			},
			Payment: Payment{
				Transaction:  order.Transaction,
				RequestID:    order.RequestID,
				Currency:     order.Currency,
				Provider:     order.Provider,
				Amount:       order.Amount,
				PaymentDt:    order.PaymentDt,
				Bank:         order.Bank,
				DeliveryCost: order.DeliveryCost,
				GoodsTotal:   order.GoodsTotal,
				CustomFee:    order.CustomFee,
			},
			Items:             resultItems,
			Locale:            order.Locale,
			InternalSignature: order.InternalSignature,
			CustomerID:        order.CustomerID,
			DeliveryService:   order.DeliveryService,
			Shardkey:          order.Shardkey,
			SmID:              order.SmID,
			DateCreated:       order.DateCreated,
			OofShard:          order.OofShard,
		})
	}

	return result, nil
}
