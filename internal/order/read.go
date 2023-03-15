package order

import (
	"github.com/fishjerky/jubo/backend/internal/order/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.FindAll()})
}

func Read(c *gin.Context) {
	// 從URL中獲取ID
	id := c.Param("id")

	// 找對應的order
	for _, p := range model.FindAll() {
		if p.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": p})
			return
		}
	}

	// 404
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
