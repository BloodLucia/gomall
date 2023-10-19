package data

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kalougata/gomall/pkg/config"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type Data struct {
	DB  *xorm.Engine
	RDB *redis.Client
}

func NewData(conf *config.Config) (*Data, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		conf.DB.User,
		conf.DB.Passwd,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.DbName,
	)
	db, err := xorm.NewEngine(conf.DB.Driver, dsn)

	if err != nil {
		return nil, nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RDB.Host, conf.RDB.Port),
		Username: conf.RDB.User,
		Password: conf.RDB.Passwd,
		DB:       conf.RDB.DB,
	})

	if err := rdb.Ping(context.Background()); err != nil {
		return nil, nil, err.Err()
	}

	data := &Data{
		DB:  db,
		RDB: rdb,
	}

	return data, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("falied to close database: %s", err)
		}
		if err := rdb.Close(); err != nil {
			log.Fatalf("falied to close redis client: %s", err)
		}
	}, nil
}
