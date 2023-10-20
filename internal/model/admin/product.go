package adminmodel

import "time"

type Product struct {
	ID             int       `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt      time.Time `xorm:"deleted TIMESTAMP deleted_at"`
	ProdName       string    `xorm:"not null VARCHAR(255) prod_name"`
	ProdIntro      string    `xorm:"not null prod_intro"`
	ProdCategoryId int       `xorm:"not null BIGINT(20) prod_category_id"`
	ProdPrice      int       `xorm:"not null DECIMAL(2) prod_price"`
	ProdStock      int       `xorm:"not null prod_stock"`
	ProdCreator    int       `xorm:"not null BIGINT(20) prod_creator"`
}

type AddProductParams struct {
	ProdName  string
	ProdIntro string
	Caregory  string
}

func (Product) TableName() string {
	return "t_mall_products"
}
