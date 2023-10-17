package adminctrl

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
)

type userController struct {
	service adminsrv.UserService
}

// Login 管理员登录
func (ctrl *userController) Login(ctx *gin.Context) {
	var reqBody adminmodel.UserLoginRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithStatusJSON(422, err.Error())
		return
	}

	v := validate.Struct(&reqBody)
	if !v.Validate() {
		ctx.AbortWithStatusJSON(400, v.Errors)
		return
	}

	resp, err := ctrl.service.Login(ctx, &reqBody)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

// Register 管理员账号注册
func (ctrl *userController) Register(ctx *gin.Context) {
	var reqBody adminmodel.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(400, err)
		ctx.Abort()
		return
	}
	v := validate.Struct(&reqBody)
	if !v.Validate() {
		ctx.JSON(400, v.Errors)
		ctx.Abort()
		return
	}

	if err := ctrl.service.Register(ctx, &reqBody); err != nil {
		log.Panic(err)
	}

	ctx.JSON(200, "ok!")
}

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewUserController(service adminsrv.UserService) UserController {
	return &userController{service}
}
