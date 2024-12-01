package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetOrder(c *gin.Context) {
	value, ok := h.cache[c.Param("orderUid")]
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, value)
}
