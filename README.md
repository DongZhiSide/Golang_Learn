# moby

[moby](https://github.com/moby/moby) 是 Docker 创建的一个开源项目，旨在支持和加速软件容器化。

它提供了一套“乐高”工具包组件、将它们组装成基于容器的定制系统的框架，以及所有容器爱好者和专业人士进行实验和交流想法的场所。组件包括容器构建工具、容器注册表、编排工具、运行时等，这些可以用作与其他工具和项目结合使用的构建块。

moby 源代码的 client 模块是 docker 的 sdk：docker 命令使用该包与守护进程通信。它也可以被你自己的 Go 应用用来完成命令行界面的功能，比如运行容器、拉取或推送图片等。

其中 api 模块是用来定义命令行客户端与守护进程通信的 HTTP API。

## moby 的参考资料

1. [Go 每日一库之121：moby（操作docker容器）](https://cloud.tencent.com/developer/article/2334511)
2. [Go client for the Docker Engine API go doc](https://pkg.go.dev/github.com/moby/moby/client)
