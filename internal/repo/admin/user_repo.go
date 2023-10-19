package adminrepo

import (
	"context"
	"fmt"

	"github.com/kalougata/gomall/internal/data"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	"github.com/kalougata/gomall/pkg/errors"
)

type userRepo struct {
	*data.Data
}

// Update implements UserRepo.
func (repo *userRepo) Update(ctx context.Context, model *adminmodel.User) error {
	if _, err := repo.DB.Context(ctx).Update(model); err != nil {
		return errors.InternalServer().WithError(err)
	}

	return nil
}

// FindByEmail implements UserRepo.
func (repo *userRepo) FindByEmail(ctx context.Context, email string) (result *adminmodel.User, has bool, err error) {
	result = &adminmodel.User{}
	has, err = repo.DB.Context(ctx).Where("email = ?", email).Get(result)
	if err != nil {
		err = errors.InternalServer().WithError(err)
	}

	return
}

// FindById 根据ID查找用户.
func (repo *userRepo) FindById(ctx context.Context, userId int) (result *adminmodel.User, has bool, err error) {
	result = &adminmodel.User{}
	has, err = repo.DB.Context(ctx).Where("id = ?", userId).Get(result)
	if err != nil {
		err = errors.InternalServer().WithError(err)
	}

	return nil, false, nil
}

// FindByLoginName 根据LoginName查找用户
func (repo *userRepo) FindByLoginName(ctx context.Context, loginName string) (result *adminmodel.User, has bool, err error) {
	result = &adminmodel.User{}
	has, err = repo.DB.Context(ctx).Where("login_name = ?", loginName).Get(result)
	if err != nil {
		fmt.Println(err)
		err = errors.InternalServer().WithError(err)
	}

	return
}

// Create 创建一个用户
func (repo *userRepo) Create(ctx context.Context, model *adminmodel.User) error {
	if _, err := repo.DB.Context(ctx).Insert(model); err != nil {
		return errors.InternalServer().WithError(err)
	}

	return nil
}

type UserRepo interface {
	Create(ctx context.Context, model *adminmodel.User) error
	FindByLoginName(ctx context.Context, loginName string) (result *adminmodel.User, has bool, err error)
	FindByEmail(ctx context.Context, email string) (result *adminmodel.User, has bool, err error)
	FindById(ctx context.Context, userId int) (result *adminmodel.User, has bool, err error)
	Update(ctx context.Context, model *adminmodel.User) error
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{Data: data}
}
