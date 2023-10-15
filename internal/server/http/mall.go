package serverhttp

import (
	"github.com/gin-gonic/gin"
)

type MallServerHTTP *gin.Engine

func NewMallServerHTTP() MallServerHTTP {
	r := gin.Default()

	mallGroup := r.Group("/api/v1/mall")
	mallGroup.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "mall")
	})

	return r
}
