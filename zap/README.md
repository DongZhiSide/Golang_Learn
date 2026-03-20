# zap

[zap](https://github.com/uber-go/zap): Uber 开源的高性能, 结构化日志库. Zap 包提供快速, 结构化, 分级的日志功能.

对于那些在热点路径中需要频繁记录日志的应用来说, 基于反射的序列化和字符串格式化代价过高, 它们会消耗大量 CPU 并产生许多小的内存分配. 换句话说, 如果你用`json.Marshal`和`fmt.Fprintf`来记录大量`interface{}`, 你的应用程序会变慢.

Zap 采用了不同的方式, 它包含一个无反射, 零分配的 JSON 编码器, 而基础的`Logger`尽可能避免序列化开销和内存分配.

在此基础上构建的高层级`SugaredLogger`, 让用户可以在需要精确控制每一次分配时使用高性能的 API, 而在更偏好熟悉, 宽松类型的 API 时也能灵活选择.

## zap 的参考资料

1. [Go 每日一库之 zap.](https://segmentfault.com/a/1190000022461706)
2. [zap go doc.](https://pkg.go.dev/go.uber.org/zap)
3. [Golang 中文学习文档 Zap.](https://golang.halfiisland.com/community/pkgs/logs/Zap.html)