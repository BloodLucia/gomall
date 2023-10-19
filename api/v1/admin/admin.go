package adminapi

import (
	"github.com/google/wire"
	adminctrl "github.com/kalougata/gomall/internal/controller/admin"
)

type AdminAPIRouter struct {
	PingCtrl adminctrl.AdminPingController
	UserCtrl adminctrl.UserController
}

func NewAdminAPIRouter(
	PingCtrl adminctrl.AdminPingController,
	UserCtrl adminctrl.UserController,
) *AdminAPIRouter {
	return &AdminAPIRouter{
		PingCtrl,
		UserCtrl,
	}
}

var AdminAPIProvider = wire.NewSet(
	adminctrl.NewPingController,
	adminctrl.NewUserController,
	NewAdminAPIRouter,
)
