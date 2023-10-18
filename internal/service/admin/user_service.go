package adminsrv

import (
	"context"
	"log"
	"time"

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
}

type UserService interface {
	Login(ctx context.Context, req *adminmodel.UserLoginRequest) (*adminmodel.UserLoginResp, error)
	Register(ctx context.Context, req *adminmodel.UserRegisterRequest) error
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

	if !utils.BcryptCheck(req.Passwd, u.Passwd) {
		return nil, errors.BadRequest("账号或密码错误")
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
	if err = srv.repo.Create(ctx, model); err != nil {
		return err
	}

	return nil
}

func NewUserService(repo adminrepo.UserRepo, jwt *jwt.JWT) UserService {
	return &userService{repo, jwt}
}
