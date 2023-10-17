package adminrepo

import (
	"context"
	"log"

	"github.com/kalougata/gomall/internal/data"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	"github.com/kalougata/gomall/pkg/e"
)

type userRepo struct {
	*data.Data
}

// FindByEmail implements UserRepo.
func (repo *userRepo) FindByEmail(ctx context.Context, email string) (result *adminmodel.User, has bool, err error) {
	result = &adminmodel.User{}
	has, err = repo.DB.Context(ctx).Where("email = ?", email).Get(result)
	if err != nil {
		log.Println(err)
		err = e.InternalServer().WithErr(err)
	}

	return
}

// FindByLoginName implements UserRepo.
func (repo *userRepo) FindByLoginName(ctx context.Context, loginName string) (result *adminmodel.User, has bool, err error) {
	result = &adminmodel.User{}
	has, err = repo.DB.Context(ctx).Where("login_name = ?", loginName).Get(result)
	if err != nil {
		log.Println(err)
		err = e.InternalServer().WithErr(err)
	}

	return
}

// Create implements UserRepo.
func (repo *userRepo) Create(ctx context.Context, model *adminmodel.User) error {
	if _, err := repo.DB.Context(ctx).Insert(model); err != nil {
		log.Println(err)
		return e.InternalServer().WithErr(err)
	}

	return nil
}

type UserRepo interface {
	Create(ctx context.Context, model *adminmodel.User) error
	FindByLoginName(ctx context.Context, loginName string) (result *adminmodel.User, has bool, err error)
	FindByEmail(ctx context.Context, email string) (result *adminmodel.User, has bool, err error)
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{Data: data}
}
