package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	server *gin.Engine
	di     *Di
}

func Init(di *Di) {
	app := new(App)

	//設定di
	app.di = di

	// 初始化GIN引擎
	app.server = gin.Default()
	app.server.Use(cors.Default()) //allow all origins

	//routing
	setApiRouter(app)

	// 啟動服務器
	app.server.Run(":8080")

}
