# zerolog

[zerolog](https://github.com/rs/zerolog) 包提供了一款轻量易用的日志库, **专门用于输出 JSON 格式日志**. zerolog 的 API 设计兼顾了优秀的开发体验与卓越性能. 其独特的链式调用 API 能避免内存分配与反射操作, 实现 JSON(或 CBOR)格式日志事件的写入.

Uber 的 zap 库是该技术方案的先驱.

zerolog 则在此基础上更进一步, 提供了更简洁易用的 API 与更出色的性能表现. 为保持代码库与 API 的简洁性, zerolog 仅专注于高效的结构化日志功能. 如需在控制台输出易读的格式化日志, 可使用库中提供的`zerolog.ConsoleWriter`(该方式性能相对较低).

## zerolog 参考资料

1. [zerolog go doc.](https://pkg.go.dev/github.com/rs/zerolog#readme-getting-started)
2. [Go 每日一库之 zerolog.](https://darjun.github.io/2020/04/24/godailylib/zerolog/)