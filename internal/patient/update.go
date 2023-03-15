package patient

import (
	"html/template"
	"net/http"

	"github.com/fishjerky/jubo-backend-api/internal/patient/model"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	// 從URL中獲取ID
	id := c.Param("id")

	// 找對應的patient
	patients := model.FindAll()
	for i, p := range patients {
		if p.ID == id {
			var updatedPatient model.Patient
			if err := c.ShouldBindJSON(&updatedPatient); err != nil {
				// 更新
				updatedPatient.ID = p.ID

				// 避免XSS
				updatedPatient.Name = template.HTMLEscapeString(updatedPatient.Name)
				patients[i] = updatedPatient

				c.JSON(http.StatusOK, gin.H{"data": updatedPatient})
				return
			}
		}
	}

	// 404
	c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
}
