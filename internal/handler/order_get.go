package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetOrder(c *gin.Context) {
	c.JSON(http.StatusOK, h.cache[c.Param("orderUid")])
}
