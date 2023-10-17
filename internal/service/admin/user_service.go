package adminsrv

import (
	"context"
	"errors"

	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
	"github.com/kalougata/gomall/pkg/utils"
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
		return errors.New("账号已被注册, 请重新输入")
	}
	if err != nil {
		return err
	}
	model := &adminmodel.User{
		LoginName: req.LoginName,
		PasswdMd5: utils.BcryptHash(req.Passwd),
	}
	if err = srv.repo.Create(ctx, model); err != nil {
		return err
	}

	return nil
}

func NewUserService(repo adminrepo.UserRepo) UserService {
	return &userService{repo}
}
