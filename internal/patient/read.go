package patient

import (
	"fmt"
	"net/http"

	"github.com/fishjerky/jubo-backend-api/internal/patient/model"
	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.FindAll()})
}

func Read(c *gin.Context) {
	// 從URL中獲取ID
	id := c.Param("id")

	// 未接DB, 需查找對應的patient
	for _, p := range model.FindAll() {
		if fmt.Sprintf("%v", p.ID) == id {
			c.JSON(http.StatusOK, gin.H{"data": p})
			return
		}
	}

	// 如果沒有找到對應的patient，返回404
	c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
}
