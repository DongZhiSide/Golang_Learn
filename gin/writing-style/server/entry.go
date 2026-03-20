package server

import "github.com/gin-gonic/gin"

var RegisterFuncs []func(router *gin.RouterGroup)

func Register(fs ...func(router *gin.RouterGroup)) {
	RegisterFuncs = append(RegisterFuncs, fs...) // maybe use copy
}

func Entry(engine *gin.Engine) {
	group := engine.Group("/style")

	for i := range RegisterFuncs {
		RegisterFuncs[i](group)
	}
}
