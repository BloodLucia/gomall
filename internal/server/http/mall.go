package serverhttp

import (
	"github.com/gin-gonic/gin"
	mallapi "github.com/kalougata/gomall/api/v1/mall"
)

type MallServerHTTP *gin.Engine

func NewMallServerHTTP(
	mar *mallapi.MallAPIRouter,
) MallServerHTTP {
	r := gin.Default()

	mallGroup := r.Group("/api/v1/mall")
	mallGroup.GET("/ping", mar.PingCtrl.Ping)

	return r
}
