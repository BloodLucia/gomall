package adminsrv

import "context"

type userService struct {
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

func NewUserService() UserService {
	return &userService{}
}
