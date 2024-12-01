package handler

import (
	"WBL0/internal/services/orders"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cache  map[string]OrderResponse
	router *gin.Engine
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.GET("/getModel/:orderUid", h.GetOrder)
}

func New(allOrders []orders.OrderResponse) *Handler {
	cache := make(map[string]OrderResponse)
	for _, order := range allOrders {
		var items []Item
		for _, item := range order.Items {
			items = append(items, Item(item))
		}
		cache[order.OrderUID] = OrderResponse{
			OrderUID:          order.OrderUID,
			TrackNumber:       order.TrackNumber,
			Entry:             order.Entry,
			Delivery:          Delivery(order.Delivery),
			Payment:           Payment(order.Payment),
			Items:             items,
			Locale:            order.Locale,
			InternalSignature: order.InternalSignature,
			CustomerID:        order.CustomerID,
			DeliveryService:   order.DeliveryService,
			Shardkey:          order.Shardkey,
			SmID:              order.SmID,
			DateCreated:       order.DateCreated,
			OofShard:          order.OofShard,
		}
	}

	h := &Handler{
		router: gin.New(),
		cache:  cache,
	}

	h.initRoutes()

	return h
}

func (h *Handler) AddModelToCache(order orders.OrderResponse) {
	var items []Item
	for _, item := range order.Items {
		items = append(items, Item(item))
	}

	h.cache[order.OrderUID] = OrderResponse{
		OrderUID:          order.OrderUID,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Delivery:          Delivery(order.Delivery),
		Payment:           Payment(order.Payment),
		Items:             items,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmID:              order.SmID,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	}
}
