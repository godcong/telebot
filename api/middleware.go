package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (a *api) middleware(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.Next()
		return
	}
	auth := context.GetHeader("Authorization")
	auths := strings.Split(auth, " ")
	if len(auths) != 2 {
		context.JSON(401, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		context.Abort()
		return
	}
	if auths[0] != "Bearer" {
		context.JSON(401, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		context.Abort()
		return
	}
	token := auths[1]
	if strings.Compare(a.config.Auth, token) == 0 {
		context.Next()
		return
	}
	context.JSON(401, gin.H{
		"code": 401,
		"msg":  "Unauthorized",
	})
}
