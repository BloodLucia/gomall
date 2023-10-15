package server

import (
	serverhttp "github.com/kalougata/gomall/internal/server/http"
)

type Server struct {
	MallServerHTTP  serverhttp.MallServerHTTP
	AdminServerHTTP serverhttp.AdminServerHTTP
}

func NewServer(msh serverhttp.MallServerHTTP, ash serverhttp.AdminServerHTTP) *Server {
	return &Server{MallServerHTTP: msh, AdminServerHTTP: ash}
}
