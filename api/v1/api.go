package v1

import (
	"github.com/google/wire"
	adminapi "github.com/kalougata/gomall/api/v1/admin"
	mallapi "github.com/kalougata/gomall/api/v1/mall"
)

var APIV1Provider = wire.NewSet(adminapi.NewAdminAPIRouter, mallapi.NewMallAPIRouter)
