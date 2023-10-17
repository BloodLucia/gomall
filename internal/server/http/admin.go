package serverhttp

import (
	"github.com/gin-gonic/gin"
	adminapi "github.com/kalougata/gomall/api/v1/admin"
)

type AdminServerHTTP *gin.Engine

func NewAdminServerHTTP(
	aar *adminapi.AdminAPIRouter,
) AdminServerHTTP {
	r := gin.Default()

	adminGroup := r.Group("/api/v1/admin")
	aar.RegisterGuestAPIRouter(adminGroup)
	return r
}
