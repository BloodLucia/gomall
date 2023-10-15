package serverhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminServerHTTP *gin.Engine

func NewAdminServerHTTP() AdminServerHTTP {
	r := gin.Default()

	adminGroup := r.Group("/api/v1/admin")
	adminGroup.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "admin")
	})
	return r
}
