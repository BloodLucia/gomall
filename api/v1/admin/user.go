package adminapi

import "github.com/gin-gonic/gin"

func (api *AdminAPIRouter) RegisterUserAPIRouter(r *gin.RouterGroup) {
	r.GET("userInfo", api.UserCtrl.GetUserInfo)
}
