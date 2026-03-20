package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 定义账号密码
	adminAccounts := gin.Accounts{"admin": "1234567"}

	// /admin 路由使用 BasicAuth 中间件
	adminAuthorized := engine.Group("/admin", gin.BasicAuthForRealm(adminAccounts, "admin need auth"))
	adminAuthorized.GET("/dashboard", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string) // 获取用户名
		c.JSON(200, gin.H{"message": "Hello " + user})
	})

	userAccounts := gin.Accounts{"user": "password"}
	userAuthorized := engine.Group("/user", gin.BasicAuthForRealm(userAccounts, "user need auth"))
	userAuthorized.GET("/dashboard", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string) // 获取用户名
		c.JSON(200, gin.H{"message": "Hello " + user})
	})
	userAuthorized.GET("/ping", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string) // 获取用户名
		c.JSON(200, gin.H{"message": "pong " + user})
	})

	engine.Run(":8080")
}
