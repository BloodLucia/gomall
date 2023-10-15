package adminapi

import "github.com/gin-gonic/gin"

// RegisterGuestAPIRouter 注册不需要登录的路由
func (api *AdminAPIRouter) RegisterGuestAPIRouter(r *gin.RouterGroup) {
	r.GET("/ping", api.PingCtrl.Ping)
	r.POST("/login", api.UserCtrl.Login)
	r.POST("/register", api.UserCtrl.Register)
}
