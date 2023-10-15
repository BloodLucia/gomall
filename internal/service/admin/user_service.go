package adminsrv

import (
	"context"

	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
)

type userService struct {
	repo adminrepo.UserRepo
}

type UserService interface {
	Login(ctx context.Context) error
	Register(ctx context.Context) error
}

// Login implements UserService.
func (srv *userService) Login(ctx context.Context) error {
	panic("unimplemented")
}

// Register implements UserService.
func (srv *userService) Register(ctx context.Context) error {
	panic("unimplemented")
}

func NewUserService(repo adminrepo.UserRepo) UserService {
	return &userService{repo}
}
