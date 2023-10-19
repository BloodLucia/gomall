package serverhttp

import (
	"github.com/gin-gonic/gin"
	adminapi "github.com/kalougata/gomall/api/v1/admin"
	"github.com/kalougata/gomall/pkg/middleware"
)

type AdminServerHTTP *gin.Engine

func NewAdminServerHTTP(
	aar *adminapi.AdminAPIRouter,
	jm *middleware.JWTMiddleware,
) AdminServerHTTP {
	r := gin.Default()
	adminGroup := r.Group("/api/v1/admin")

	// 需要登录的路由
	authGroup := adminGroup.Group("")
	authGroup.Use(jm.AdminJWT())
	aar.RegisterUserAPIRouter(authGroup)

	// 不需要登录的路由
	noAuthGroup := adminGroup.Group("")
	aar.RegisterGuestAPIRouter(noAuthGroup)

	return r
}
