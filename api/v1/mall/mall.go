package mallapi

import (
	"github.com/google/wire"
	mallctrl "github.com/kalougata/gomall/internal/controller/mall"
)

type MallAPIRouter struct {
	PingCtrl mallctrl.MallPingController
}

func NewMallAPIRouter(
	PingCtrl mallctrl.MallPingController,
) *MallAPIRouter {
	return &MallAPIRouter{PingCtrl}
}

var MallAPIProvider = wire.NewSet(
	mallctrl.NewPingController,
	NewMallAPIRouter,
)
