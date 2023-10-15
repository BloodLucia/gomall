package adminapi

import "github.com/gin-gonic/gin"

// RegisterGuestAPIRouter 注册不需要登录的路由
func (api *AdminAPIRouter) RegisterGuestAPIRouter(r *gin.RouterGroup) {
	r.POST("/login")
	r.POST("/register")
}
