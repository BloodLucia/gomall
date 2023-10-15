package mallctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingController struct {
}

// Ping implements PingController.
func (*pingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "mall,pong!",
	})
}

type MallPingController interface {
	Ping(ctx *gin.Context)
}

func NewPingController() MallPingController {
	return &pingController{}
}
