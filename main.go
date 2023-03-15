package main

import (
	"github.com/fishjerky/jubo/backend/app"
	"github.com/fishjerky/jubo/backend/configs"
)

func main() {
	// load config
	config := configs.Init()

	// setup di
	di := app.InitDi(config)

	// app init & start
	app.Init(di)
}
