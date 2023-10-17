package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalougata/gomall/pkg/e"
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Build(ctx *gin.Context, err *e.Error, data any) *response {
	if err == nil {
		return &response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: data,
		}
	}

	return &response{
		Code: err.Code,
		Msg:  err.Msg,
		Data: nil,
	}
}
