package order

import (
	"fmt"
	"net/http"

	"github.com/fishjerky/jubo-backend-api/internal/order/model"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	// 從URL中獲取ID
	id := c.Param("id")

	// 未接DB, 需查找對應的patient
	orders := model.FindAll()
	for i, p := range orders {
		fmt.Printf("Patient: %s Order: %s\n", p.ID, id)
		if p.ID == id {
			var updatedOrder model.Order
			fmt.Printf("Update Order: %s Order: %s\n", p.ID, id)
			if err := c.ShouldBindJSON(&updatedOrder); err != nil {
				fmt.Printf("Update Order Fail. Error: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return

			} else {
				updatedOrder.ID = p.ID
				orders[i] = updatedOrder

				// 返回更新後的patient
				c.JSON(http.StatusOK, gin.H{"data": updatedOrder})
				return
			}
		}
	}

	// 404
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
