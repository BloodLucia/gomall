//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	adminapi "github.com/kalougata/gomall/api/v1/admin"
	mallapi "github.com/kalougata/gomall/api/v1/mall"
	"github.com/kalougata/gomall/internal/server"
	serverhttp "github.com/kalougata/gomall/internal/server/http"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		mallapi.MallAPIProvider,
		adminapi.AdminAPIProvider,
		serverhttp.NewAdminServerHTTP,
		serverhttp.NewMallServerHTTP,
		server.NewServer,
	))
}
