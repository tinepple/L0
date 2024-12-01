package handler

import (
	"WBL0/internal/services/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cache  cache.Cache
	router *gin.Engine
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.GET("/getModel/:orderUid", h.GetOrder)
}

func New(cache cache.Cache) *Handler {
	h := &Handler{
		router: gin.New(),
		cache:  cache,
	}

	h.initRoutes()

	return h
}
