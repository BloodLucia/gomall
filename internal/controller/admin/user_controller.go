package adminctrl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
	"github.com/kalougata/gomall/pkg/errors"
	"github.com/kalougata/gomall/pkg/response"
	"github.com/kalougata/gomall/pkg/validator"
)

type userController struct {
	service adminsrv.UserService
}

// Login 管理员登录
func (ctrl *userController) Login(ctx *gin.Context) {
	var reqBody adminmodel.UserLoginRequest
	if errs := validator.Checker(ctx, reqBody); errs != nil {
		fmt.Println(errs)
		response.Build(ctx, errors.UnprocessableEntity(), errs)
		return
	}

	if res, err := ctrl.service.Login(ctx, &reqBody); err == nil {
		response.Build(ctx, nil, res)
	} else {
		response.Build(ctx, err, nil)
		return
	}
}

// Register 管理员账号注册
func (ctrl *userController) Register(ctx *gin.Context) {
	var reqBody adminmodel.UserRegisterRequest
	if errs := validator.Checker(ctx, reqBody); errs != nil {
		response.Build(ctx, errors.UnprocessableEntity(), errs)
		return
	}

	if err := ctrl.service.Register(ctx, &reqBody); err != nil {
		response.Build(ctx, err, nil)
		return
	}

	response.Build(ctx, nil, nil)
}

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewUserController(service adminsrv.UserService) UserController {
	return &userController{service}
}
