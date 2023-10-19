package configs

import "github.com/kelseyhightower/envconfig"

type Redis struct {
	Host   string `required:"true"`
	Port   int    `required:"true"`
	DB     int    `required:"true"`
	User   string
	Passwd string
}

func RedisConfig() Redis {
	var rdb Redis
	envconfig.MustProcess("REDIS", &rdb)
	return rdb
}
