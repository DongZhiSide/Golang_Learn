# viper

[viper](https://github.com/spf13/viper) 是适用于 Go 应用(包括十二因素应用)的完整配置解决方案. 它可嵌入任意应用, 能处理各类配置需求与格式. 其支持功能如下:

- 设置默认值, 显式值.
- 读取配置文件, 跨多个位置动态发现配置文件, 读取远程系统配置(如 Etcd 或 Consul).
- 读取环境变量.
- 读取命令行标志, 读取缓冲区数据.
- 实时监听并更新配置.
- 为配置键设置别名, 方便重构.
- Viper 可被视为承载应用所有配置需求的注册中心.

注意: 在使用`viper.Unmarshal`函数前, 一定要保证结构体的字段是**可导出**的, 否则无法解析到结构体上.

## viper 的参考资料

1. [viper readme.](https://github.com/spf13/viper/blob/master/README.md)
2. [Golang 中文学习文档-viper.](https://golang.halfiisland.com/community/pkgs/config/Viper.html)

> 参考 1 提到 viper 支持 ini 等格式, 但实际测试发现, 新版本的 viper 不支持 ini 格式了, 见此 [issue](https://github.com/spf13/viper/issues/2092).
