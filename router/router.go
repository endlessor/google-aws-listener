package router

import (
	"github.com/gin-gonic/gin"
	"google-rtb/api"
	"google-rtb/config"
)

func GetRouter() *gin.Engine {
	r := gin.New()

	r.GET("/", api.StatusCheck)
	r.POST("api/rtb", api.RtbListener)

	return r
}

func GetPort() string {
	return config.Cfg.ServerConfigurations.Port
}
