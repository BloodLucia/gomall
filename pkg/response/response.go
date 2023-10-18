package response

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	myErrors "github.com/kalougata/gomall/pkg/errors"
)

type respBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Build(ctx *gin.Context, err error, data any) {
	if err == nil {
		ctx.JSON(http.StatusOK, respBody{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: data,
		})
		return
	}

	var myErr *myErrors.Error

	if !errors.As(err, &myErr) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respBody{
			Code: http.StatusInternalServerError,
			Msg:  "未知错误",
			Data: nil,
		})
	}

	if myErrors.IsInternalServer(myErr) {
		log.Println(myErr)
	}

	ctx.AbortWithStatusJSON(myErr.Code, respBody{
		Code: myErr.Code,
		Msg:  myErr.Message,
		Data: data,
	})
}
