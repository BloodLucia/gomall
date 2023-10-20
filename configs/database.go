package configs

import "github.com/kelseyhightower/envconfig"

type Database struct {
	Driver string `required:"true"`
	Host   string `default:"127.0.0.1"`
	Port   int    `default:"3306"`
	User   string `required:"true"`
	Passwd string `required:"true"`
	DbName string `required:"true"`
}

func DatabaseConfig() Database {
	var db Database
	envconfig.MustProcess("DB", &db)
	return db
}
