package adminrepo

import (
	"context"
	"errors"
	"log"

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
		return nil, false, err
	}
	if !has {
		return nil, false, nil
	}

	return result, true, nil
}

// Create implements UserRepo.
func (repo *userRepo) Create(ctx context.Context, model *adminmodel.User) error {
	_, has, err := repo.FindByLoginName(ctx, model.LoginName)
	if err != nil {
		// TODO
		log.Panicf("查询数据库出错: %s", err)
	}
	if has {
		return errors.New("登录名已被使用, 请重新输入")
	}

	if _, err := repo.DB.Context(ctx).Insert(model); err != nil {
		// TODO
		log.Panicf("插入数据库出错: %s", err)
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
