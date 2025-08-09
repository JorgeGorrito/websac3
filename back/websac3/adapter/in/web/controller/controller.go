package controller

import (
	"websac3/common/jwt"

	"github.com/gin-gonic/gin"
)

type Authenticable struct{}

func (c *Authenticable) GetToken(ctx *gin.Context) *jwt.AccessTokenClaims {
	var claims *jwt.AccessTokenClaims
	if value, exists := ctx.Get("token"); exists {
		if typedValue, ok := value.(*jwt.AccessTokenClaims); ok {
			claims = typedValue
		}
	}
	return claims
}
