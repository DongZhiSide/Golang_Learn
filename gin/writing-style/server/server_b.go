package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Register(ServerBRouter)
}

func ServerBRouter(group *gin.RouterGroup) {
	group.GET("/pingb", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pongb"}) })
}
