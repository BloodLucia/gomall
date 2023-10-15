package data

import (
	"context"
	"log"

	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData() (*Data, func(), error) {
	db, err := xorm.NewEngine("", "")

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
		if err := db.Close();err!=nil {
			log.Fatalf("falied to close database: %s", err)
		}
	}, nil
}
