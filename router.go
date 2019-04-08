package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func NewRouter(token string) *gin.Engine {
	engine := gin.Default()

	engine.Any("/ping", func(ctx *gin.Context) {
		log.Println("ping")
		ctx.JSON(http.StatusOK, "pong")

	})
	engine.Any("/"+token, func(ctx *gin.Context) {
		bytes, e := ioutil.ReadAll(ctx.Request.Body)
		log.Println(string(bytes), e)
		ctx.JSON(http.StatusOK, "ok")
	})
	return engine
}
