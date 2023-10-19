package validator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/kalougata/gomall/pkg/errors"
	"github.com/kalougata/gomall/pkg/response"
)

func Checker[T any](ctx *gin.Context, value T) validate.Errors {
	if err := ctx.ShouldBindJSON(&value); err != nil {
		response.Build(ctx, errors.New(http.StatusUnprocessableEntity, "UnproleEntity"), nil)
		return nil
	}

	v := validate.Struct(&value)
	if !v.Validate() {
		return v.Errors
	}

	return nil
}
