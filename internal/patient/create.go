package patient

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/fishjerky/jubo-backend-api/internal/patient/model"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// parameter
	var newPatient model.Patient
	if err := c.ShouldBindJSON(&newPatient); err != nil {
		// 避免XSS
		newPatient.Name = template.HTMLEscapeString(newPatient.Name)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// new ID
	newID := len(model.FindAll()) + 1
	newPatient.ID = strconv.Itoa(newID)

	// insert
	model.Insert(newPatient)

	// response
	c.JSON(http.StatusOK, gin.H{"data": newPatient})
}
