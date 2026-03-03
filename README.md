# snowflake

[snowflake](https://github.com/bwmarrin/snowflake): 一个非常简单的推特雪花生成器. 解析现有雪花 ID 的方法. 将雪花 ID 转换为其他几种数据类型以及转换回来的方法. 函数可以方便地在 JSON API 中使用雪花 ID. 单调时钟计算可防止时钟漂移.

默认情况下, ID 格式遵循最初的 Twitter 雪花格式. 整个 ID 是一个 63 位整数, 存储在`int64`中, 41 位用于以毫秒精度存储时间戳, 采用自定义纪元. 10 位用于存储节点 ID, 范围从 0 到 1023. 12 位用于存储序列号, 范围从 0 到 4095. 你可以通过设置雪花来更改节点 ID 和步数(序列)所用的位数.

每次生成 ID 时就像这样: 以毫秒精度存储时间戳, 使用 41 位的 ID 存储. 然后在后续的位中加入 NodeID. 然后添加序列号, 从 0 开始, 随着每个 ID 在同一毫秒内生成, 逐一递增. 如果你在同一毫秒内生成足够多的 ID, 导致序列翻滚或过填, 生成函数就会暂停到下一个毫秒.

## snowflake 的参考资料

1. [snowflake go doc.](https://pkg.go.dev/github.com/bwmarrin/snowflake)
