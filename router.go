package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type router struct {
}

func NewRouter() {
	engine := gin.Default()

	engine.Any("/ping", func(context *gin.Context) {
		log.Println("ping")
	})

}
