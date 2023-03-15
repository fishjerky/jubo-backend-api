package patient

import (
	"fmt"
	"net/http"

	"github.com/fishjerky/jubo-backend-api/internal/patient/model"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	// 從URL中獲取ID
	id := c.Param("id")

	// 遍歷patient數組，查找對應的patient
	patients := model.FindAll()
	for i, p := range patients {
		if fmt.Sprintf("%v", p.ID) == id {
			// 從數組中刪除patient
			patients = append(patients[:i], patients[i+1:]...)

			// 返回成功的响应
			c.JSON(http.StatusOK, gin.H{"data": true})
			return
		}
	}

	// 如果沒有找到對應的patient，返回404
	c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
}
