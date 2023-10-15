package adminapi

import (
	"github.com/google/wire"
	adminctrl "github.com/kalougata/gomall/internal/controller/admin"
)

type AdminAPIRouter struct {
	PingCtrl adminctrl.AdminPingController
}

func NewAdminAPIRouter(PingCtrl adminctrl.AdminPingController) *AdminAPIRouter {
	return &AdminAPIRouter{PingCtrl}
}

var AdminAPIProvider = wire.NewSet(
	adminctrl.NewPingController,
	NewAdminAPIRouter,
)
