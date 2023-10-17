package data

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData() (*Data, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		"",
		"",
		"",
		3306,
		"gomall_dev",
	)
	db, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		return nil, nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, nil, err
	}

	data := &Data{
		DB: db,
	}

	return data, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("falied to close database: %s", err)
		}
	}, nil
}
