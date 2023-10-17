package adminsrv

import (
	"context"

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
	model := &adminmodel.User{
		LoginName: req.LoginName,
		PasswdMd5: req.Passwd,
	}
	return srv.repo.Create(ctx, model)
}

func NewUserService(repo adminrepo.UserRepo) UserService {
	return &userService{repo}
}
