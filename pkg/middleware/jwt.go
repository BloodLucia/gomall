package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt/v5"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
	myErrs "github.com/kalougata/gomall/pkg/errors"
	"github.com/kalougata/gomall/pkg/jwt"
	"github.com/kalougata/gomall/pkg/response"
	"github.com/spf13/cast"
)

type JWTMiddleware struct {
	aus adminsrv.UserService
	aur adminrepo.UserRepo
	jwt *jwt.JWT
}

func NewJWTMiddleware(
	aus adminsrv.UserService,
	aur adminrepo.UserRepo,
	jwt *jwt.JWT,
) *JWTMiddleware {
	return &JWTMiddleware{aus, aur, jwt}
}

func (jm *JWTMiddleware) AdminJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			response.Build(ctx, myErrs.Unauthorized(), nil)
			return
		}
		claims, err := jm.jwt.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwtPkg.ErrTokenExpired) {
				response.Build(ctx, myErrs.Unauthorized().WithMsg("token已过期"), nil)
			}
			response.Build(ctx, myErrs.Unauthorized().WithMsg("token校验失败"), nil)
			return
		}
		if claims.UserRule == "" || claims.UserRule != "admin" {
			response.Build(ctx, myErrs.Forbidden(), nil)
			return
		}

		if user, has, err := jm.aur.FindById(ctx, cast.ToInt(claims.UserId)); err == nil && has {
			ctx.Set("user", user)
			ctx.Next()
		} else {
			response.Build(ctx, myErrs.Unauthorized(), nil)
			return
		}
	}
}

func (jm *JWTMiddleware) MallJWT(j *jwt.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.Abort()
			return
		}
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
