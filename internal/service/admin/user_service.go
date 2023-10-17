package adminsrv

import (
	"context"
	"errors"
	"log"

	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
)

type userService struct {
	repo adminrepo.UserRepo
}

type UserService interface {
	Login(ctx context.Context) error
	Register(ctx context.Context, req *adminmodel.UserRegisterRequest) error
}

// Login implements UserService.
func (srv *userService) Login(ctx context.Context) error {
	panic("unimplemented")
}

// Register implements UserService.
func (srv *userService) Register(ctx context.Context, req *adminmodel.UserRegisterRequest) error {
	_, has, err := srv.repo.FindByLoginName(ctx, req.LoginName)
	if has {
		return errors.New("用户名已被注册, 请重新输入")
	}
	if err != nil {
		log.Panicf("查询数据库出错: %s", err)
	}
	model := &adminmodel.User{
		LoginName: req.LoginName,
		PasswdMd5: req.Passwd,
	}
	if err = srv.repo.Create(ctx, model); err != nil {
		log.Panicf("插入数据库出错: %s", err)
	}

	return nil
}

func NewUserService(repo adminrepo.UserRepo) UserService {
	return &userService{repo}
}
