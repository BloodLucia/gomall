package adminctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingController struct {
}

// Ping implements PingController.
func (*pingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "admin,pong!",
	})
}

type AdminPingController interface {
	Ping(ctx *gin.Context)
}

func NewPingController() AdminPingController {
	return &pingController{}
}
