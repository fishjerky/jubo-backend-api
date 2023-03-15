package app

import (
	"github.com/fishjerky/jubo/backend/internal/order"
	"github.com/fishjerky/jubo/backend/internal/patient"
)

func setApiRouter(app *App) {

	// patient routing
	app.server.GET("/patients", patient.ReadAll)
	app.server.GET("/patients/:id", patient.Read)
	app.server.POST("/patients", patient.Create)
	app.server.PUT("/patients/:id", patient.Update)
	app.server.DELETE("/patients/:id", patient.Delete)

	//order routing
	app.server.GET("/orders", order.ReadAll)
	app.server.GET("/order/:id", order.Read)
	app.server.PUT("/order/:id", order.Update)
}
