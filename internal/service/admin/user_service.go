package adminsrv

import (
	"context"
	"log"
	"time"

	"github.com/kalougata/gomall/internal/data"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
	"github.com/kalougata/gomall/pkg/errors"
	"github.com/kalougata/gomall/pkg/jwt"
	"github.com/kalougata/gomall/pkg/utils"
	"github.com/spf13/cast"
)

type userService struct {
	repo adminrepo.UserRepo
	jwt  *jwt.JWT
	data *data.Data
}

type UserService interface {
	Login(ctx context.Context, req *adminmodel.UserLoginRequest) (*adminmodel.UserLoginResp, error)
	Register(ctx context.Context, req *adminmodel.UserRegisterRequest) error
	GetUserInfo(ctx context.Context, userId int) (userInfo *adminmodel.UserInfo, err error)
	UpdateUserInfo(ctx context.Context, req *adminmodel.UpdateUserInfoRequest) error
}

// UpdateUserInfo 修改管理员信息
func (srv *userService) UpdateUserInfo(ctx context.Context, req *adminmodel.UpdateUserInfoRequest) error {
	user, has, err := srv.repo.FindById(ctx, req.ID)
	if err != nil {
		return err
	}
	if !has {
		return errors.NotFound("用户不存在")
	}

	user.LoginName = req.LoginName
	user.NickName = req.NickName
	user.Email = req.Email

	if err := srv.repo.Update(ctx, user); err != nil {
		return errors.InternalServer().WithMsg("更新用户信息失败, 请稍后再试")
	}

	return nil
}

// GetUserInfo 获取管理员信息
func (srv *userService) GetUserInfo(ctx context.Context, userId int) (userInfo *adminmodel.UserInfo, err error) {
	if user, has, err := srv.repo.FindById(ctx, userId); err == nil && has {
		userInfo = &adminmodel.UserInfo{
			ID:        user.ID,
			LoginName: user.LoginName,
			NickName:  user.NickName,
			Locked:    user.Locked,
			CreatedAt: user.CreatedAt,
		}

		return userInfo, nil
	} else {
		return nil, errors.InternalServer().WithMsg("获取用户信息失败")
	}
}

// Login implements UserService.
func (srv *userService) Login(ctx context.Context, req *adminmodel.UserLoginRequest) (*adminmodel.UserLoginResp, error) {
	u, has, err := srv.repo.FindByLoginName(ctx, req.LoginName)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.NotFound("账号不存在")
	}
	claims := jwt.MyCustomClaims{
		UserId:    cast.ToString(u.ID),
		LoginName: u.LoginName,
		UserRule:  "admin",
	}

	token, err := srv.jwt.BuildToken(claims, time.Now().Add(time.Minute*10))

	if err != nil {
		log.Printf("failed to create token: %s \n", err)
		return nil, errors.InternalServer().WithMsg("创建Token失败")
	}

	return &adminmodel.UserLoginResp{Token: token}, nil
}

// Register implements UserService.
func (srv *userService) Register(ctx context.Context, req *adminmodel.UserRegisterRequest) error {
	_, has, err := srv.repo.FindByLoginName(ctx, req.LoginName)
	if has {
		return errors.BadRequest("账号已被注册, 请重新输入")
	}
	if err != nil {
		return err
	}
	model := &adminmodel.User{
		LoginName: req.LoginName,
		Passwd:    utils.BcryptHash(req.Passwd),
	}
	if err := srv.repo.Create(ctx, model); err != nil {
		return errors.InternalServer().WithMsg("注册账号失败, 请稍后重试")
	}

	return nil
}

func NewUserService(repo adminrepo.UserRepo, jwt *jwt.JWT, data *data.Data) UserService {
	return &userService{repo, jwt, data}
}
