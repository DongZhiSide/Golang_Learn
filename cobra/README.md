# cobra

[Cobra](https://github.com/spf13/cobra): Cobra 提供简洁的接口, 用于创建功能强大的现代命令行界面(CLI), 风格与 git 和 go 工具一脉相承.

```shell title="Cobra 推荐使用的目录结构"
appName/
    cmd/
        add.go
        your.go
        commands.go
        here.go
    main.go
```

Cobra 有一个很大的缺点, 它的短标志(`-`开头的标志)只能是一个字符, 多了不行.

## Cobra 的参考资料

1. [Cobra go doc.](https://pkg.go.dev/github.com/spf13/cobra#section-readme)
2. [Cobra user guide.](https://github.com/spf13/cobra/blob/v1.10.2/site/content/user_guide.md)(重点参考)

## Cobra 的概念与术语

- 命令: 操作单元. 为了支持子命令, Cobra 的整个命令结构是一个树形结构, 从根命令出发到叶子命令, 每个节点都是一个命令.
- 参数: 命令后跟随的"位置参数"(无`-`或`--`前缀), 是命令的"必选或可选"输入.
- 标志: 命令的"可选配置项"(有`-`或`--`前缀), 是命令的"必选或可选"配置项.
- 持久性和本地性(Persistent and Local): 即父命令的部分配置会继承到子命令中. 例如, 父命令的`--ip`标志, 如果该标志无持久性(则有本地性, 互斥性质), 那么子命令无法使用`--ip`标志; 如果该标志有持久性, 那么子命令可以使用`--ip`标志.
- Run: 解析完参数和标志后要执行的函数, 这也是我们要实现的函数. 有 5 个 Run 系列的函数, 其调用顺序是`PersistentPreRun`, `PreRun`, `Run`, `PostRun`, `PersistentPostRun()`. 只有 Run 存在, 才会 PreRun 和 PostRun 系列函数.

> Run 系列还一个 RunE 系列的函数, 会返回一个 error, 如果 Run 系列和 RunE 系列都存在, 则只会执行 RunE 系列.
