package order

import (
	"fmt"
	"github.com/fishjerky/jubo/backend/internal/order/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	// 從URL中獲取ID
	id := c.Param("id")

	orders := model.FindAll()
	for i, p := range orders {
		fmt.Printf("Patient: %s Order: %s\n", p.ID, id)
		if p.ID == id {
			var updatedOrder model.Order
			fmt.Printf("Update Order: %s Order: %s\n", p.ID, id)
			if err := c.ShouldBindJSON(&updatedOrder); err != nil {
				fmt.Printf("Update Order Fail. Error: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return // 返回錯誤響應，並結束函數

			} else {
				updatedOrder.ID = p.ID
				orders[i] = updatedOrder

				// 返回更新後的patient
				c.JSON(http.StatusOK, gin.H{"data": updatedOrder})
				return // 更新後結束函數
			}
		}
	}

	// 如果沒有找到對應的patient，返回404
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
