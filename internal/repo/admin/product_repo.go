package adminrepo

import (
	"context"

	"github.com/kalougata/gomall/internal/data"
	adminmodel "github.com/kalougata/gomall/internal/model/admin"
)

type ProductRepo interface {
	Create(ctx context.Context, model *adminmodel.Product) error
}

type productRepo struct {
	data *data.Data
}

// AddProduct implements ProductRepo.
func (repo *productRepo) Create(ctx context.Context, model *adminmodel.Product) error {
	if _, err := repo.data.DB.Context(ctx).Insert(model); err != nil {
		return err
	}

	return nil
}

func NewProductRepo(data *data.Data) ProductRepo {
	return &productRepo{data}
}
