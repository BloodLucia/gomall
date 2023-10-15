//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	adminapi "github.com/kalougata/gomall/api/v1/admin"
	mallapi "github.com/kalougata/gomall/api/v1/mall"
	"github.com/kalougata/gomall/internal/data"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
	"github.com/kalougata/gomall/internal/server"
	serverhttp "github.com/kalougata/gomall/internal/server/http"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		mallapi.MallAPIProvider,
		adminapi.AdminAPIProvider,
		data.NewData,
		adminrepo.AdminRepoProvider,
		adminsrv.AdminServiceProvider,
		serverhttp.NewAdminServerHTTP,
		serverhttp.NewMallServerHTTP,
		server.NewServer,
	))
}
