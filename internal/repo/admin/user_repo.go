package adminrepo

import (
	"context"
	"errors"

	"github.com/kalougata/gomall/internal/data"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
)

type userRepo struct {
	*data.Data
}

// FindByLoginName implements UserRepo.
func (repo *userRepo) FindByLoginName(ctx context.Context, loginName string) (result *adminmodel.User, has bool, err error) {
	result = &adminmodel.User{}
	has, err = repo.DB.Context(ctx).Where("login_name = ?", loginName).Get(result)
	if err != nil {
		err = errors.New("查询数据库失败")
	}

	return
}

// Create implements UserRepo.
func (repo *userRepo) Create(ctx context.Context, model *adminmodel.User) error {
	if _, err := repo.DB.Context(ctx).Insert(model); err != nil {
		return errors.New("插入数据库失败")
	}

	return nil
}

type UserRepo interface {
	Create(ctx context.Context, model *adminmodel.User) error
	FindByLoginName(ctx context.Context, loginName string) (result *adminmodel.User, has bool, err error)
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{Data: data}
}
