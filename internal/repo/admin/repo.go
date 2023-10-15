package adminrepo

import "github.com/google/wire"

var AdminRepoProvider = wire.NewSet(
	NewUserRepo,
)
