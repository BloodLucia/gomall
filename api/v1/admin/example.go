package adminapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (aar *AdminAPIRouter) RegisterExampleRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong!",
			"server": "admin",
		})
	})
}
