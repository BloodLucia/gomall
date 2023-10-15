//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kalougata/gomall/internal/server"
	serverhttp "github.com/kalougata/gomall/internal/server/http"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		serverhttp.NewAdminServerHTTP,
		serverhttp.NewMallServerHTTP,
		server.NewServer,
	))
}
