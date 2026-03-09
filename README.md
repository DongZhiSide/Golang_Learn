# Golang_Learn

[grpc](https://github.com/grpc/grpc-go): 由 go 实现的 gRPC: 一个高性能, 开源的通用 RPC 框架，将移动端和 HTTP/2 置于首位.

## grpc 的参考资料

1. [grpc 官方, Introduction to gRPC.](https://grpc.io/docs/what-is-grpc/introduction/)
2. [Protocol Buffers Documentation.](https://protobuf.dev/overview/)
3. [grpc 的 github 官方仓库示例代码.](https://github.com/grpc/grpc-go/tree/master/examples)
4. [李文周的博客, gRPC 教程.](https://www.liwenzhou.com/posts/Go/gRPC/)

## rpc 概念

远程过程调用(remote procedure call, rpc)指计算机程序在当前进程的不同地址空间(通常是在共享计算机网络的另一台计算机上)执行一个过程(子程序), 该过程以正常(本地)过程调用的形式写入, 程序员无需明确写入远程交互的详细信息.

简单来说: 客户端在不知道调用细节的情况下, 调用存在于远程计算机上的某个对象或者函数, 就像调用本地应用程序中的对象或者函数一样.

### 为什么要有 rpc

其实这是应用开发到一定的阶段的强烈需求驱动的. 如果我们开发简单的单一应用, 逻辑简单, 用户不多, 流量不大, 那我们用不着. 当我们的系统访问量增大, 业务增多时, 我们会发现一台单机运行此系统已经无法承受.

此时, 我们可以将业务拆分成几个互不关联的应用, 分别部署在各自机器上, 以划清逻辑并减小压力. 此时, 我们也可以不需要RPC, 因为应用之间是互不关联的.

当我们的业务越来越多, 应用也越来越多时, 自然的, 我们会发现有些功能已经不能简单划分开来或者划分不出来. 此时, 可以将公共业务逻辑抽离出来, 将之组成独立的服务 Service 应用.

而原有的, 新增的应用都可以与那些独立的 Service 应用交互, 以此来完成完整的业务功能. 所以此时, 我们急需一种**高效的应用程序之间的通讯手段**来完成这种需求, 所以你看, RPC 大显身手的时候来了!

其实描述的场景也是服务化, 微服务和分布式系统架构的基础场景. 即 RPC 框架就是实现以上结构的有力方式.

> 虽然这些 Service 应用可以使用一些传统的通讯协议来交换数据, 例如 http, 但是这些 Service 应用都要重复的编写有关于 http 的样板代码, 例如路由注册, 参数绑定, 错误处理等.
>
> 因此我们希望有工具能通过代码生成的方式去生成这些样板代码. 这些工具为了提高性能, 不一定会去使用常见的数据传输格式, 例如 json, 而是自己设计一套数据传输格式.
>
> 这些工具为了泛用性, 又会专门设计一套定义 api 接口的语言, 称接口定义语言(IDL), 这样, 这些工具通过解析 IDL, 就能生成不同编程语言的代码出来.
>
> 这些工具也就是 rpc 框架.

### 参考

1. [Remote procedure call.](https://en.wikipedia.org/wiki/Remote_procedure_call)
2. [RPC框架: 从原理到选型, 一文带你搞懂RPC. ](https://blog.csdn.net/daobuxinzi/article/details/133931185)
