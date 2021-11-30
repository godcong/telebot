package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (a *api) middleware(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.Next()
		return
	}
	auth := ctx.GetHeader("Authorization")
	auths := strings.Split(auth, " ")
	if len(auths) != 2 {
		ctx.JSON(401, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		ctx.Abort()
		return
	}
	if auths[0] != "Bearer" {
		ctx.JSON(401, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		ctx.Abort()
		return
	}
	token := auths[1]
	if strings.Compare(a.config.Auth, token) == 0 {
		ctx.Next()
		return
	}
	ctx.JSON(401, gin.H{
		"code": 401,
		"msg":  "Unauthorized",
	})
}
