package adminctrl

import (
	"github.com/gin-gonic/gin"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
)

type userController struct {
	service adminsrv.UserService
}

// Login 管理员登录
func (*userController) Login(ctx *gin.Context) {
	ctx.String(200, "login")
}

// Register 管理员账号注册
func (*userController) Register(ctx *gin.Context) {
	ctx.String(200, "register")
}

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewUserController(service adminsrv.UserService) UserController {
	return &userController{service}
}
