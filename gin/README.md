# gin

[Gin](https://github.com/gin-gonic/gin): Gin 是一个高性能的 HTTP Web 框架, 用 Go 语言编写. 它提供了类似 Martini 的 API, 但性能显著更好——得益于 httprouter, 速度最高可提升至 40 倍. Gin 专为构建 REST API, Web 应用和微服务而设计.

Gin 将 Express.js 风格的简洁路由与 Go 的高性能特性结合在一起, 非常适合:

- 构建高吞吐量的 REST API.
- 开发需要处理大量并发请求的微服务.
- 创建对响应速度要求极高的 Web 应用.
- 快速原型化 Web 服务, 减少样板代码.

Gin 的核心特性:

- 零分配路由器: 极度节省内存的路由机制, 无堆分配.
- 高性能: 基准测试显示其速度优于其他 Go Web 框架.
- 中间件支持: 可扩展的中间件系统, 用于认证, 日志, CORS 等.
- 防崩溃: 内置恢复中间件, 防止`panic`导致服务器崩溃.
- JSON 验证: 自动绑定和验证请求/响应的 JSON 数据.
- 路由分组: 组织相关路由并应用通用中间件.
- 错误管理: 集中化的错误处理与日志记录.
- 内置渲染: 支持 JSON, XML, HTML 模板等多种输出格式.
- 可扩展性.

## gin 的参考资料

1. [Gin 官方文档.](https://gin-gonic.com/zh-cn/docs/)

## gin 编码结构参考

参考自: [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin).

```shell
.
cmd         # 存放 main 函数
config      # 存放配置文件映射成结构体的代码
initialize  # 存放初始化代码, 如初始化路由等
api         # 存放 gin 的控制器代码
router      # 存放 gin 的路由注册代码
service     # 存放业务逻辑代码
model       # 存放数据库模型代码
utils       # 存放业务无关代码
```

```go
// 这是常见的 gin 控制器
func Ctrl(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// 有些写法写成了这样:
// 这是为了让 gin 控制器具有参数配置功能
func BuildCtrl() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
    }
}

// 例如: 在创建控制器时能根据传入参数来回复不同的内容
// 当然也可以使用结构体来传递参数, 但这可能会污染结构体的"空间"
func BuildCtrl(msg string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{"message": msg})
    }
}
router.GET("/ping", BuildCtrl("pong"))

// go 没有命名空间, 对于不同模块但是类似的函数,
// 可以把函数变成方法以避免冲突
type api1 struct{}
func (api1) Ping(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

type api2 struct{}
func (api2) Ping(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
```

最后, 目录结构与你写的系统架构有关, 没有绝对的对错, 只要合理即可.
