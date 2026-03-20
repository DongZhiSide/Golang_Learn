package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Register(ServerARouter)
}

func ServerARouter(group *gin.RouterGroup) {
	group.GET("/pinga", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "ponga"}) })
}
