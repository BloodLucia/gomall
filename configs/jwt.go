package configs

import "github.com/kelseyhightower/envconfig"

type JWT struct {
	Key string `required:"true"`
}

func JWTStore() JWT {
	var jwt JWT
	envconfig.MustProcess("JWT", &jwt)
	return jwt
}
