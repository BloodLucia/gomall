package middleware

import (
	"github.com/gin-gonic/gin"
	adminrepo "github.com/kalougata/gomall/internal/repo/admin"
	adminsrv "github.com/kalougata/gomall/internal/service/admin"
	"github.com/kalougata/gomall/pkg/jwt"
)

type JWTMiddleware struct {
	aus adminsrv.UserService
	aur adminrepo.UserRepo
}

func NewJWTMiddleware(
	aus adminsrv.UserService,
	aur adminrepo.UserRepo,
) *JWTMiddleware {
	return &JWTMiddleware{aus, aur}
}

func (jm *JWTMiddleware) AdminJWT(j *jwt.JWT) gin.HandlerFunc {
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
		if claims.UserRule == "" || claims.UserRule != "admin" {
			ctx.Abort()
			return
		}

		if user, has, err := jm.aur.FindByLoginName(ctx, claims.LoginName); err == nil && has {
			ctx.Set("claims", claims)
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.Abort()
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
