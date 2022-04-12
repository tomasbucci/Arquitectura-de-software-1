package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.Run(":3000")

}

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}
