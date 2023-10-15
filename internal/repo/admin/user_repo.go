package adminrepo

import (
	"context"

	"github.com/kalougata/gomall/internal/data"
)

type userRepo struct {
	*data.Data
}

// Create implements UserRepo.
func (repo *userRepo) Create(ctx context.Context) error {
	panic("unimplemented")
}

type UserRepo interface {
	Create(ctx context.Context) error
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{Data: data}
}
