package main

import (
	"github.com/chenzr/gin-web/config"
	"github.com/chenzr/gin-web/logger"
	"github.com/chenzr/gin-web/model"
	"github.com/chenzr/gin-web/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	cfg := config.GetConfig()
	gin.SetMode(cfg.Mode)
	logger.GetLogger().Traceln(viper.AllSettings())
	g := gin.Default()
	router.Load(g)
	model.Load()
	_ = g.Run(":" + cfg.Port)   
}
