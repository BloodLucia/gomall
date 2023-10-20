package adminsrv

import (
	"context"

	adminmodel "github.com/kalougata/gomall/internal/model/admin"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
	"github.com/kalougata/gomall/pkg/errors"
)

type productService struct {
	repo adminrepo.ProductRepo
}

type ProductService interface {
	AddProduct(ctx context.Context, req *adminmodel.AddProductParams) error
}

// AddProduct implements ProductService.
func (srv *productService) AddProduct(ctx context.Context, req *adminmodel.AddProductParams) error {
	product := &adminmodel.Product{
		ProdName: req.ProdName,
	}
	if err := srv.repo.Create(ctx, product); err != nil {
		return errors.InternalServer().WithMsg("添加商品失败, 请稍后再试")
	}

	return nil
}

func NewProductService(repo adminrepo.ProductRepo) ProductService {
	return &productService{repo}
}
