package adminctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
	"github.com/kalougata/gomall/pkg/errors"
	"github.com/kalougata/gomall/pkg/response"
)

type userController struct {
	service adminsrv.UserService
}

// UpdateUserPasswd 更新管理员信息
func (ctrl *userController) UpdateUserInfo(ctx *gin.Context) {
	var reqBody adminmodel.UpdateUserInfoRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Build(ctx, errors.UnprocessableEntity(), err.Error())
		return
	}

	v := validate.Struct(reqBody)
	if !v.Validate() {
		response.Build(ctx, errors.UnprocessableEntity(), v.Errors)
		return
	}

	if err := ctrl.service.UpdateUserInfo(ctx, &reqBody); err != nil {
		response.Build(ctx, err, nil)
		return
	}

	response.Build(ctx, nil, nil)
}

// GetUserInfo 获取管理员的信息
func (ctrl *userController) GetUserInfo(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		response.Build(ctx, nil, nil)
		return
	}

	userId := value.(*adminmodel.User).ID
	if user, err := ctrl.service.GetUserInfo(ctx, userId); err != nil {
		response.Build(ctx, err, nil)
		return
	} else {
		response.Build(ctx, nil, user)
	}
}

// Login 管理员登录
func (ctrl *userController) Login(ctx *gin.Context) {
	var reqBody adminmodel.UserLoginRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Build(ctx, errors.UnprocessableEntity(), err)
		return
	}

	v := validate.Struct(&reqBody)
	if !v.Validate() {
		response.Build(ctx, errors.UnprocessableEntity(), v.Errors)
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
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Build(ctx, errors.UnprocessableEntity(), err.Error())
		return
	}
	v := validate.Struct(&reqBody)
	if !v.Validate() {
		response.Build(ctx, errors.UnprocessableEntity(), v.Errors)
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
	GetUserInfo(ctx *gin.Context)
	UpdateUserInfo(ctx *gin.Context)
}

func NewUserController(service adminsrv.UserService) UserController {
	return &userController{service}
}
