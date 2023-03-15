package app

import "github.com/fishjerky/jubo/backend/configs"

type Di struct {
	config *configs.Config
	//logger *zap.Logger
}

var di = new(Di)

func InitDi(config *configs.Config) *Di {
	di.config = config
	return di
}
